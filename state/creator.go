// Copyright (c) 2018 The VeChainThor developers
// Copyright (c) 2019 The PlayMaker developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package state

import (
	"github.com/playmakerchain/powerplay/kv"
	"github.com/playmakerchain/powerplay/powerplay"
)

// Creator state creator to cut-off kv dependency.
type Creator struct {
	kv kv.GetPutter
}

// NewCreator create a new state creator.
func NewCreator(kv kv.GetPutter) *Creator {
	return &Creator{kv}
}

// NewState create a new state object.
func (c *Creator) NewState(root powerplay.Bytes32) (*State, error) {
	return New(root, c.kv)
}
