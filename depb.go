package depb

import (
	"fmt"

	cmn "github.com/rhuairahrighairidh/go-mod-test-common"
)

func GetVersion() string {
	return fmt.Sprintf("depB: %s, common: %s", "v0.1.0", cmn.GetVersion())
}
