// Copyright (c) 2018 The VeChainThor developers
// Copyright (c) 2019 The PlayMaker developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package chain

import (
	"github.com/playmakerchain/powerplay/block"
	"github.com/playmakerchain/powerplay/powerplay"
)

// Seeker to seek block by given number on the chain defined by head block ID.
type Seeker struct {
	chain       *Chain
	headBlockID powerplay.Bytes32
	err         error
}

func newSeeker(chain *Chain, headBlockID powerplay.Bytes32) *Seeker {
	return &Seeker{
		chain:       chain,
		headBlockID: headBlockID,
	}
}

func (s *Seeker) setError(err error) {
	if s.err == nil {
		s.err = err
	}
}

// Err returns error occurred.
func (s *Seeker) Err() error {
	return s.err
}

// GetID returns block ID by the given number.
func (s *Seeker) GetID(num uint32) powerplay.Bytes32 {
	if num > block.Number(s.headBlockID) {
		panic("num exceeds head block")
	}
	id, err := s.chain.GetAncestorBlockID(s.headBlockID, num)
	s.setError(err)
	return id
}

// GetHeader returns block header by the given number.
func (s *Seeker) GetHeader(id powerplay.Bytes32) *block.Header {
	header, err := s.chain.GetBlockHeader(id)
	if err != nil {
		s.setError(err)
		return &block.Header{}
	}
	return header
}

// GenesisID get genesis block ID.
func (s *Seeker) GenesisID() powerplay.Bytes32 {
	return s.chain.GenesisBlock().Header().ID()
}
