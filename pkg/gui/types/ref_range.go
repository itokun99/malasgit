package types

import "github.com/itokun99/malasgit/pkg/commands/models"

type RefRange struct {
	From models.Ref
	To   models.Ref
}
