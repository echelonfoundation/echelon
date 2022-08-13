package keeper

import (
	"encoding/hex"
	"encoding/binary"
	"github.com/coniks-sys/coniks-go/crypto/vrf"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
	"strconv"
)

func (k Keeper) CreateRandomNumber(ctx sdk.Context, msg *types.MsgCreateRandom) error {

	userval, isFound := k.GetUserval(ctx, msg.Creator)

	var user_key_count int64 = 1
	if isFound {
		user_key_count = userval.Count + 1
	}

	sk, err := vrf.GenerateKey(nil)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Secret Key is not generated")
	}

	random_val_key := msg.Creator + "," + strconv.FormatInt(user_key_count, 10)
	a_message := []byte(random_val_key)

	vrv, proof := sk.Prove(a_message) // Generate vrv (verifiable random value) and proof
	pub_key, ok_bool := sk.Public()   // public key creation

	if ok_bool == false {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Public Key is not generated")
	}

	var max_val_uint64 uint64 = 18446744073709551615
	parse_vrv_to_uint64 := binary.BigEndian.Uint64(vrv)
	var float_vrv float64 = float64(parse_vrv_to_uint64) / float64(max_val_uint64)
	final_vrv := float_vrv * float64(msg.Multiplier)
	final_vrv_float := float_vrv * float64(msg.Multiplier)

	newRandomVal := types.Randomval{
		Index:     random_val_key,
		Creator:   msg.Creator,
		Vrv:       hex.EncodeToString(vrv),
		Multiplier:msg.Multiplier,
		Proof:     hex.EncodeToString(proof),
		Pubk:      hex.EncodeToString(pub_key),
		Message:   random_val_key,
		Parsedvrv: binary.BigEndian.Uint64(vrv),
		Floatvrv:  float_vrv,
		Finalvrv:  uint64(final_vrv),
		Finalvrvfl: final_vrv_float,
	}

	newUserVal := types.Userval{
		Index:    msg.Creator,
		Useraddr: msg.Creator,
		Count:    user_key_count,
	}

	k.SetRandomval(ctx, newRandomVal)
	k.SetUserval(ctx, newUserVal)
	return nil
}

func (k Keeper) VerifyRandomNumber(ctx sdk.Context, req *types.QueryVerifyValuesRequest) (string, error) {

	var public_key vrf.PublicKey
	public_key, err := hex.DecodeString(req.Pubkey)
	if err != nil {
		return "false", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Public Key cannot be decoded")
	}
	message_value := []byte(req.Message)
	vrv_value, err := hex.DecodeString(req.Vrv)
	if err != nil {
		return "false", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "VRV Value cannot be decoded")
	}

	proof_value, err := hex.DecodeString(req.Proof)
	if err != nil {
		return "false", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Proof VRF value cannot be decoded")
	}

	is_verified := public_key.Verify(message_value, vrv_value, proof_value)

	return strconv.FormatBool(is_verified), err

}
