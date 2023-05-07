package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furychain/fury/x/house/types"
)

// Deposit performs deposit operation to participate as a house in a specific market/order book
func (k msgServer) Deposit(goCtx context.Context,
	msg *types.MsgDeposit,
) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if err := msg.ValidateSanity(ctx, &params); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid deposit")
	}

	var payload types.DepositTicketPayload
	if err := k.ovmKeeper.VerifyTicketUnmarshal(goCtx, msg.Ticket, &payload); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketVerification, "%s", err)
	}

	if err := payload.Validate(msg.Creator); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketPayloadValidation, "%s", err)
	}

	participationIndex, err := k.Keeper.Deposit(ctx, msg.Creator, msg.MarketUID, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to deposit")
	}

	return &types.MsgDepositResponse{
		MarketUID:          msg.MarketUID,
		ParticipationIndex: participationIndex,
	}, nil
}
