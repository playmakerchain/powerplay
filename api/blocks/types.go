// Copyright (c) 2018 The VeChainThor developers
// Copyright (c) 2019 The PlayMaker developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package blocks

import (
	"github.com/playmakerchain//block"
	"github.com/playmakerchain//"
)

//Block block
type Block struct {
	Number       uint32         `json:"number"`
	ID           .Bytes32   `json:"id"`
	Size         uint32         `json:"size"`
	ParentID     .Bytes32   `json:"parentID"`
	Timestamp    uint64         `json:"timestamp"`
	GasLimit     uint64         `json:"gasLimit"`
	Beneficiary  .Address   `json:"beneficiary"`
	GasUsed      uint64         `json:"gasUsed"`
	TotalScore   uint64         `json:"totalScore"`
	TxsRoot      .Bytes32   `json:"txsRoot"`
	StateRoot    .Bytes32   `json:"stateRoot"`
	ReceiptsRoot .Bytes32   `json:"receiptsRoot"`
	Signer       .Address   `json:"signer"`
	IsTrunk      bool           `json:"isTrunk"`
	Transactions [].Bytes32 `json:"transactions"`
}

func convertBlock(b *block.Block, isTrunk bool) (*Block, error) {
	if b == nil {
		return nil, nil
	}
	signer, err := b.Header().Signer()
	if err != nil {
		return nil, err
	}
	txs := b.Transactions()
	txIds := make([].Bytes32, len(txs))
	for i, tx := range txs {
		txIds[i] = tx.ID()
	}

	header := b.Header()
	return &Block{
		Number:       header.Number(),
		ID:           header.ID(),
		ParentID:     header.ParentID(),
		Timestamp:    header.Timestamp(),
		TotalScore:   header.TotalScore(),
		GasLimit:     header.GasLimit(),
		GasUsed:      header.GasUsed(),
		Beneficiary:  header.Beneficiary(),
		Signer:       signer,
		Size:         uint32(b.Size()),
		StateRoot:    header.StateRoot(),
		ReceiptsRoot: header.ReceiptsRoot(),
		TxsRoot:      header.TxsRoot(),
		IsTrunk:      isTrunk,
		Transactions: txIds,
	}, nil
}
