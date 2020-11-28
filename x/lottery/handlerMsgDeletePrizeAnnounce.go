package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ppalineee/lottery/x/lottery/types"
	"github.com/ppalineee/lottery/x/lottery/keeper"
)

// Handle a message to delete name
func handleMsgDeletePrizeAnnounce(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeletePrizeAnnounce) (*sdk.Result, error) {
	if !k.PrizeAnnounceExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetPrizeAnnounceOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeletePrizeAnnounce(ctx, msg.ID)
	return &sdk.Result{}, nil
}
