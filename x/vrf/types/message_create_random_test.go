package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/echelonfoundation/echelon/v3/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRandom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRandom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRandom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRandom{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
