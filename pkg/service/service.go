package service

import (
	"context"
	"github.com/g3co/twtest/pkg/ethrpc"
	"github.com/g3co/twtest/pkg/tools"
	"log"
	"strings"
	"time"
)

const pullingInterval = time.Second * 1

//go:generate mockgen -source=service.go -destination=mock/service_mock.go
type storageProvider interface {
	GetCurrentBlock() (string, error)
	AddAddress(address string) error
	SaveTX(tx ethrpc.Transaction) error
	GetTXByAddress(address string) ([]ethrpc.Transaction, error)
}

type chainViewerProvider interface {
	GetCurrentBlock() (string, error)
	GetBlockInfo(blockNumber string) (*ethrpc.Block, error)
}

type Service struct {
	storage storageProvider
	viewer  chainViewerProvider
}

func NewService(
	storage storageProvider,
	viewer chainViewerProvider,
) *Service {
	return &Service{
		storage: storage,
		viewer:  viewer,
	}
}

func (s *Service) Run(ctx context.Context) error {
	ticker := time.NewTicker(pullingInterval)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := s.processBlock(); err != nil {
				log.Println(err)
			}
		}
	}
}

func (s *Service) processBlock() error {
	chainBlockNum, err := s.viewer.GetCurrentBlock()
	if err != nil {
		return err
	}

	storageBlockNum, err := s.storage.GetCurrentBlock()
	if err != nil {
		return err
	}

	if chainBlockNum == storageBlockNum {
		// skipping, block has been processed
		return nil
	}

	block, err := s.viewer.GetBlockInfo(chainBlockNum)
	if err != nil {
		return err
	}

	log.Printf("Processing block %s (%d transactions)\n",
		chainBlockNum, len(block.Transactions))

	for _, tx := range block.Transactions {
		if txErr := s.storage.SaveTX(tx); txErr != nil {
			log.Printf("Failed to save transaction %s Err: %s\n", tx.Hash, err)
		}
	}

	return nil
}

// GetCurrentBlock last parsed block
func (s *Service) GetCurrentBlock() int {
	blockNumber, err := s.storage.GetCurrentBlock()
	if err != nil {
		log.Println("GetCurrentBlock err:", err)
		return 0
	}

	blockNumberInt, err := tools.ConvertHexToInt(blockNumber)
	if err != nil {
		log.Println("GetCurrentBlock err:", err)
		return 0
	}

	return int(blockNumberInt)
}

// Subscribe add address to observer
func (s *Service) Subscribe(address string) bool {
	if len(address) != 42 {
		return false
	}

	if !strings.HasPrefix(address, "0x") {
		return false
	}

	err := s.storage.AddAddress(address)
	if err != nil {
		log.Println("AddAddress err:", err)
		return false
	}
	return true
}

// GetTransactions list of inbound or outbound transactions for an address
func (s *Service) GetTransactions(address string) []ethrpc.Transaction {
	txs, err := s.storage.GetTXByAddress(address)
	if err != nil {
		log.Println("GetTXByAddress err:", err)
		return nil
	}
	return txs
}
