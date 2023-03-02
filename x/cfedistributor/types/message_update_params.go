package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const TypeMsgUpdateAllSubDistributors = "update_all_subdistributors"

var _ sdk.Msg = &MsgUpdateAllSubDistributors{}

func (msg *MsgUpdateAllSubDistributors) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAllSubDistributors) Type() string {
	return TypeMsgUpdateAllSubDistributors
}

func (msg *MsgUpdateAllSubDistributors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAllSubDistributors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAllSubDistributors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = Params{SubDistributors: msg.SubDistributors}.Validate()
	if err != nil {
		sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateSubDistributor = "update_single_subdistributor"

var _ sdk.Msg = &MsgUpdateSubDistributor{}

func (msg *MsgUpdateSubDistributor) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributor) Type() string {
	return TypeMsgUpdateSubDistributor
}

func (msg *MsgUpdateSubDistributor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = msg.SubDistributor.Validate()
	if err != nil {
		sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}
