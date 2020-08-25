package std

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types and vesting
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	vesting.RegisterInterfaces(interfaceRegistry)
}
