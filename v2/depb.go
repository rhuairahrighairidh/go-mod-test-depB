package depb

import (
	"fmt"

	cmn "github.com/rhuairahrighairidh/go-mod-test-common"
)

func GetThePkgVersion() string {
	// downgrade common pkg
	return fmt.Sprintf("depB (v2): %s, common: %s", "v2.0.1", cmn.GetVersion())
}
