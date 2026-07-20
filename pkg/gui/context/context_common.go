package context

import (
	"github.com/itokun99/malasgit/pkg/common"
	"github.com/itokun99/malasgit/pkg/gui/types"
)

type ContextCommon struct {
	*common.Common
	types.IGuiCommon
}
