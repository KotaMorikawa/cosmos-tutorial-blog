package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/blog module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrMissingPostTitle = sdkerrors.Register(ModuleName, 1101, "title is missing")
	ErrMissingPostBody = sdkerrors.Register(ModuleName, 1102, "body is missing")
)
