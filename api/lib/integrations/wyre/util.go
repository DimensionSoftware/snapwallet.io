package wyre

import (
	"strings"
	"time"

	proto "github.com/khoerling/flux/api/lib/protocol"
)

func WyreTransferToProto(t *Transfer) *proto.WyreTransfer {
	return &proto.WyreTransfer{
		Id:             t.ID,
		Source:         strings.Split(t.Source, ":")[1],
		Dest:           strings.Split(t.Dest, ":")[1],
		SourceCurrency: t.SourceCurrency,
		DestCurrency:   t.DestCurrency,
		SourceAmount:   t.SourceAmount,
		DestAmount:     t.DestAmount,
		Message:        t.Message,
		ExchangeRate:   t.ExchangeRate,
		Fees:           t.Fees,
		TotalFees:      t.TotalFees,
		Blockhash:      t.BlockchainTx.Blockhash,
		NetworkTxId:    t.BlockchainTx.NetworkTxID,
		Status:         t.Status,
		CreatedAt:      time.Unix(t.CreatedAt/1000, 0).Format(time.RFC3339),
		ExpiresAt:      time.Unix(t.ExpiresAt/1000, 0).Format(time.RFC3339),
		CompletedAt:    time.Unix(t.CompletedAt/1000, 0).Format(time.RFC3339),
		CancelledAt:    time.Unix(t.CancelledAt/1000, 0).Format(time.RFC3339),
	}
}
