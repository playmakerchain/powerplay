// Copyright (c) 2018 The VeChainThor developers
// Copyright (c) 2019 The PlayMaker developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package energy

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/playmakerchain/powerplay/lvldb"
	"github.com/playmakerchain/powerplay/state"
	"github.com/playmakerchain/powerplay/powerplay"
)

func TestEnergy(t *testing.T) {
	kv, _ := lvldb.NewMem()
	st, _ := state.New(powerplay.Bytes32{}, kv)

	acc := powerplay.BytesToAddress([]byte("a1"))

	eng := New(powerplay.BytesToAddress([]byte("eng")), st, 0)
	tests := []struct {
		ret      interface{}
		expected interface{}
	}{
		{eng.Get(acc), &big.Int{}},
		{func() bool { eng.Add(acc, big.NewInt(10)); return true }(), true},
		{eng.Get(acc), big.NewInt(10)},
		{eng.Sub(acc, big.NewInt(5)), true},
		{eng.Sub(acc, big.NewInt(6)), false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.ret)
	}

	assert.Nil(t, st.Err())
}

func TestEnergyGrowth(t *testing.T) {
	kv, _ := lvldb.NewMem()
	st, _ := state.New(powerplay.Bytes32{}, kv)

	acc := powerplay.BytesToAddress([]byte("a1"))

	st.SetEnergy(acc, &big.Int{}, 10)

	vetBal := big.NewInt(1e18)
	st.SetBalance(acc, vetBal)

	bal1 := New(powerplay.Address{}, st, 1000).
		Get(acc)

	x := new(big.Int).Mul(powerplay.EnergyGrowthRate, vetBal)
	x.Mul(x, new(big.Int).SetUint64(1000-10))
	x.Div(x, big.NewInt(1e18))

	assert.Equal(t, x, bal1)

	assert.Nil(t, st.Err())

}
