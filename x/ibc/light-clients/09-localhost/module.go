package localhost

import (
	"github.com/oraichain/orai/x/ibc/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
