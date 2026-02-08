package errorx

import "errors"

var (
	AuthTokenNotFound     = errors.New("auth: no token found")
	AuthSignMethod        = errors.New("auth: invalid signing method")
	AuthTokenClaims       = errors.New("auth: failed to claim token")
	AuthInvalidCredential = errors.New("auth: invalid credential")

	NoDataSaved = errors.New("no data saved")
	DataInvalid = errors.New("data invalid")
)
