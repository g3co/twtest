package tools

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertHexToInt(t *testing.T) {
	tests := []struct {
		input  string
		output int64
		err    bool
	}{
		{input: "0xa", output: 10},
		{input: "0x3", output: 3},
		{input: "0x0", output: 0},
		{input: "0", err: true},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			o, err := ConvertHexToInt(tc.input)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.output, o)
			}
		})
	}
}

func TestConvertIntToHex(t *testing.T) {
	tests := []struct {
		input  int64
		output string
	}{
		{output: "0xa", input: 10},
		{output: "0x3", input: 3},
		{output: "0x0", input: 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			o := ConvertIntToHex(tc.input)
			require.Equal(t, tc.output, o)
		})
	}
}
