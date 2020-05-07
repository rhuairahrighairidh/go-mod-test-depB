package depb

import (
	"fmt"

	cmn "github.com/rhuairahrighairidh/go-mod-test-common"
)

func GetThePkgVersion() string {
	cmn.NewFunctionality()
	return fmt.Sprintf("depB (v2): %s, common: %s", "v2.0.0", cmn.GetVersion())
}
