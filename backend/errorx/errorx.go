package errorx

import "errors"

var (
	AuthTokenNotFound     = errors.New("auth: no token found")
	AuthTokenInvalid      = errors.New("auth: invalid token")
	AuthMethodInvalid     = errors.New("auth: invalid token method")
	AuthSignMethod        = errors.New("auth: invalid signing method")
	AuthTokenClaims       = errors.New("auth: failed to claim token")
	AuthInvalidCredential = errors.New("auth: invalid credential")

	NoDataSaved   = errors.New("no data saved")
	NoDataUpdated = errors.New("no data updated")
	NoDataFound   = errors.New("no data found")
	DataInvalid   = errors.New("data invalid")
	DataExists    = errors.New("data exists")
)
