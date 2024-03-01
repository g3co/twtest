package tools

import (
	"errors"
	"fmt"
	"strconv"
)

func ConvertHexToInt(d string) (int64, error) {
	if len(d) < 3 {
		return 0, errors.New("wrong format")
	}
	return strconv.ParseInt(d[2:], 16, 64)
}

func ConvertIntToHex(d int64) string {
	return fmt.Sprintf("0x%x", d)
}
