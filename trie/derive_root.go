// Copyright (c) 2018 The VechainThor Developers
// Copyright (c) 2019 The PlayMaker developers

package trie

import (
	"bytes"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/playmakerchain/powerplay/powerplay"
)

// see "github.com/ethereum/go-ethereum/types/derive_sha.go"

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveRoot(list DerivableList) powerplay.Bytes32 {
	keybuf := new(bytes.Buffer)
	trie := new(Trie)
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		trie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	return trie.Hash()
}
