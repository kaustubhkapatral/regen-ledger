package server

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	BalancePrefix          byte = 0x0
	SupplyPrefix           byte = 0x1
	MaxDecimalPlacesPrefix byte = 0x2
	DenomAdminPrefix       byte = 0x3
)

func BalanceKey(addr sdk.AccAddress, denom string) []byte {
	key := make([]byte, len(addr)+len(denom)+2)
	addrLen := len(addr)
	if addrLen > 256 {
		panic(fmt.Errorf("address %x is too long (max 256 bytes)", addr))
	}
	key[0] = BalancePrefix
	key[1] = byte(addrLen)
	copy(key[2:2+addrLen], addr)
	copy(key[2+addrLen:], denom)
	return key
}

func SupplyKey(denom string) []byte {
	key := make([]byte, len(denom)+1)
	key[0] = SupplyPrefix
	copy(key[1:], denom)
	return key
}

func MaxDecimalPlacesKey(denom string) []byte {
	key := make([]byte, len(denom)+1)
	key[0] = MaxDecimalPlacesPrefix
	copy(key[1:], denom)
	return key
}

func DenomAdminKey(denom string) []byte {
	key := make([]byte, len(denom)+1)
	key[0] = DenomAdminPrefix
	copy(key[1:], denom)
	return key
}