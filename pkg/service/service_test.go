package service

import (
	"errors"
	"github.com/g3co/twtest/pkg/ethrpc"
	mock_service "github.com/g3co/twtest/pkg/service/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetCurrentBlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	viewer := mock_service.NewMockchainViewerProvider(ctrl)
	storage := mock_service.NewMockstorageProvider(ctrl)
	svc := NewService(storage, viewer)
	t.Run("success", func(t *testing.T) {
		storage.EXPECT().GetCurrentBlock().Return("0x12714f4", nil)
		currentBlock := svc.GetCurrentBlock()
		require.Equal(t, 19338484, currentBlock)
	})

	t.Run("error from storage", func(t *testing.T) {
		storage.EXPECT().GetCurrentBlock().Return("", errors.New("sample"))
		currentBlock := svc.GetCurrentBlock()
		require.Equal(t, 0, currentBlock)
	})

	t.Run("wrong data in storage", func(t *testing.T) {
		storage.EXPECT().GetCurrentBlock().Return("", nil)
		currentBlock := svc.GetCurrentBlock()
		require.Equal(t, 0, currentBlock)
	})
}

func TestSubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	viewer := mock_service.NewMockchainViewerProvider(ctrl)
	storage := mock_service.NewMockstorageProvider(ctrl)
	svc := NewService(storage, viewer)
	t.Run("success", func(t *testing.T) {
		storage.EXPECT().AddAddress("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5").
			Return(nil)
		ok := svc.Subscribe("0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5")
		require.True(t, ok)
	})
	t.Run("wrong address", func(t *testing.T) {
		ok := svc.Subscribe("test")
		require.False(t, ok)
	})
	t.Run("wrong address format", func(t *testing.T) {
		ok := svc.Subscribe("0095222290dd7278aa3ddd389cc1e1d165cc4bafe5")
		require.False(t, ok)
	})
}

func TestProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	viewer := mock_service.NewMockchainViewerProvider(ctrl)
	storage := mock_service.NewMockstorageProvider(ctrl)
	svc := NewService(storage, viewer)

	t.Run("success", func(t *testing.T) {
		viewer.EXPECT().GetCurrentBlock().Return("0x12714f4", nil)
		storage.EXPECT().GetCurrentBlock().Return("0x12714f3", nil)

		tx := ethrpc.Transaction{BlockNumber: "0x12714f4"}
		viewer.EXPECT().GetBlockInfo("0x12714f4").
			Return(&ethrpc.Block{
				Transactions: []ethrpc.Transaction{
					tx,
				},
			}, nil)
		storage.EXPECT().SaveTX(tx).Return(nil)
		require.NoError(t, svc.processBlock())
	})

	t.Run("already processed", func(t *testing.T) {
		viewer.EXPECT().GetCurrentBlock().Return("0x12714f4", nil)
		storage.EXPECT().GetCurrentBlock().Return("0x12714f4", nil)
		require.NoError(t, svc.processBlock())
	})
}
