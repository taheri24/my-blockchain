package handlers

import "errors"

var (
	errUnmatchedToken                  = errors.New("keys of currentSignatures and keys of nextSignatures are not same")
	errInvalidTokenBySignatureMismatch = errors.New(`format keys of currentSignatures(or nextSignatures) is "token:minerId", please check tokens or miners`)
	errInvalidTokenByFormat            = errors.New(`format keys of currentSignatures(or nextSignatures) is "number:minerId", please check key`)
	errUnexpectedSuccess               = errors.New(`it is no error but it is unexpected too`)
)
