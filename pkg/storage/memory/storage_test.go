package memory

import (
	"encoding/json"
	"github.com/g3co/twtest/pkg/ethrpc"
	"github.com/stretchr/testify/require"
	"testing"
)

const txSample = `  {
    "blockHash": "0xcf6c01ec452f73a1c3dcdfe68ba5d70da6ca0b01d2dd95afe1485bbdb11ea646",
    "blockNumber": "0x12714f4",
    "from": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5",
    "gas": "0x6ac1",
    "gasPrice": "0xaffdc638b",
    "hash": "0xb37936d9e59eeed69592ed984d46413cb5a2330f51b1b2737c0620429b69029f",
    "input": "0x",
    "nonce": "0xb1ebf",
    "to": "0x8306300ffd616049fd7e4b0354a64da835c1a81c",
    "transactionIndex": "0x6e",
    "value": "0x4c2aa381b1f343",
    "v": "0x1",
    "r": "0xcb4e1f179b96cb55a220c6391b9b60d42e664bc254d1c620e70e1591251a37c1",
    "s": "0x6c8e3e778557114ad5881ceb3b5d6d48a6ab70594cf81c0b330e8ec128f61d77"
  }`

func TestStorage(t *testing.T) {
	st := NewStorage()

	var tx ethrpc.Transaction
	err := json.Unmarshal([]byte(txSample), &tx)
	require.NoError(t, err)

	err = st.SaveTX(tx)
	require.NoError(t, err)

	cb, err := st.GetCurrentBlock()
	require.NoError(t, err)
	require.Equal(t, "0x12714f4", cb)

	var emptyTXs []ethrpc.Transaction

	txs, err := st.GetTXByAddress("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5")
	require.NoError(t, err)
	require.Equal(t, emptyTXs, txs)

	txs, err = st.GetTXByAddress("0x8306300ffd616049fd7e4b0354a64da835c1a81c")
	require.NoError(t, err)
	require.Equal(t, emptyTXs, txs)

	err = st.AddAddress("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5")
	require.NoError(t, err)

	err = st.AddAddress("0x8306300ffd616049fd7e4b0354a64da835c1a81c")
	require.NoError(t, err)

	err = st.SaveTX(tx)
	require.NoError(t, err)

	txs, err = st.GetTXByAddress("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5")
	require.NoError(t, err)
	require.Equal(t, tx, txs[0])

	txs, err = st.GetTXByAddress("0x8306300ffd616049fd7e4b0354a64da835c1a81c")
	require.NoError(t, err)
	require.Equal(t, tx, txs[0])
}
