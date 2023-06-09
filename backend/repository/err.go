package repository

import (
	"sthl/constants"
	"sthl/ent"
)

func handleEntRepoErr(err error) error {
	var rerr error
	switch {
	case ent.IsValidationError(err):
		rerr = constants.ErrBadRequest
	case ent.IsConstraintError(err):
		rerr = constants.ErrBadRequest
	case ent.IsNotFound(err):
		rerr = constants.ErrNotFound
	case ent.IsNotLoaded(err):
		rerr = constants.ErrInternalServer
	case ent.IsNotSingular(err):
		rerr = constants.ErrInternalServer
	default:
		rerr = constants.ErrInternalServer
	}
	return rerr
}
