package chainviewer

import (
	"github.com/g3co/twtest/pkg/ethrpc"
	"sync/atomic"
)

const rpcVersion = "2.0"

type Viewer struct {
	sequence atomic.Int64
}

func NewViewer() *Viewer {
	return &Viewer{}
}

func (s *Viewer) getNextSequence() int64 {
	return s.sequence.Add(1)
}

func (s *Viewer) GetCurrentBlock() (string, error) {
	var resp ethrpc.Response[string]
	err := sendRequest(ethrpc.Request{
		Jsonrpc: rpcVersion,
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      s.getNextSequence(),
	}, &resp)

	if err != nil {
		return "", err
	}

	return resp.Result, nil
}

func (s *Viewer) GetBlockInfo(blockNumber string) (*ethrpc.Block, error) {
	var resp ethrpc.Response[ethrpc.Block]
	err := sendRequest(ethrpc.Request{
		Jsonrpc: rpcVersion,
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumber, true},
		Id:      s.getNextSequence(),
	}, &resp)

	if err != nil {
		return nil, err
	}

	return &resp.Result, nil
}
