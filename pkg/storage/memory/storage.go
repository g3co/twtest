package memory

import "github.com/g3co/twtest/pkg/ethrpc"

type Storage struct {
	txStorage map[string][]ethrpc.Transaction
	lastBlock string
}

func NewStorage() *Storage {
	return &Storage{
		txStorage: map[string][]ethrpc.Transaction{},
	}
}

// GetCurrentBlock returns the last processed block.
func (s *Storage) GetCurrentBlock() (string, error) {
	return s.lastBlock, nil
}

// AddAddress saves address to the internal memory.
func (s *Storage) AddAddress(address string) error {
	if _, ok := s.txStorage[address]; !ok {
		s.txStorage[address] = []ethrpc.Transaction{}
	}

	return nil
}

// SaveTX saves transaction and bind it to related addresses.
// Address has to be added in advance via AddAddress method.
func (s *Storage) SaveTX(tx ethrpc.Transaction) error {
	if _, ok := s.txStorage[tx.From]; ok {
		s.txStorage[tx.From] = append(s.txStorage[tx.From], tx)
	}

	if _, ok := s.txStorage[tx.To]; ok {
		s.txStorage[tx.To] = append(s.txStorage[tx.To], tx)
	}

	s.lastBlock = tx.BlockNumber

	return nil
}

// GetTXByAddress returns all saved transactions for the requested address
func (s *Storage) GetTXByAddress(address string) ([]ethrpc.Transaction, error) {
	txs := s.txStorage[address]
	return txs, nil
}
