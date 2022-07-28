package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRandom{}

func NewMsgCreateRandom(creator string, multiplier uint64) *MsgCreateRandom {
	return &MsgCreateRandom{
		Creator:   creator,
		Multiplier: multiplier,
	}
}

func (msg *MsgCreateRandom) Route() string {
	return RouterKey
}

func (msg *MsgCreateRandom) Type() string {
	return "CreateRandom"
}

func (msg *MsgCreateRandom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRandom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRandom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
