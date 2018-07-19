package ethutils

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

func IsValidEthereumAddress(addr ...string) ([]string, bool) {
	invalid := make([]string, 0, len(addr))
	for _, a := range addr {
		if !common.IsHexAddress(a) && !common.IsHexAddress(fmt.Sprintf("0x%s", a)) {
			invalid = append(invalid, fmt.Sprintf("\"%s\"", a))
		}
	}

	return invalid, len(invalid) == 0
}