package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ppalineee/lottery/x/lottery/types"
	"github.com/ppalineee/lottery/x/lottery/keeper"
)

func handleMsgSetLottery(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetLottery) (*sdk.Result, error) {
	var lottery = types.Lottery{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Name: msg.Name,
    	Detail: msg.Detail,
    	Reward: msg.Reward,
    	Drawdate: msg.Drawdate,
    	Price: msg.Price,
	}
	if !msg.Creator.Equals(k.GetLotteryOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetLottery(ctx, lottery)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
