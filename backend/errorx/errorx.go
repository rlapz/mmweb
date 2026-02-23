package errorx

import "errors"

var (
	AuthTokenNotFound     = errors.New("auth: no token found")
	AuthSignMethod        = errors.New("auth: invalid signing method")
	AuthTokenClaims       = errors.New("auth: failed to claim token")
	AuthInvalidCredential = errors.New("auth: invalid credential")

	NoDataSaved   = errors.New("no data saved")
	NoDataUpdated = errors.New("no data updated")
	DataInvalid   = errors.New("data invalid")
	DataExists    = errors.New("data exists")
)
