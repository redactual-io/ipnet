package ipnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVer_String(t *testing.T) {
	tests := []struct {
		name string
		v    Ver
		want string
	}{
		{
			"v4",
			v4,
			"v4",
		},
		{
			"v6",
			v6,
			"v6",
		},
		{
			"Unknown",
			(Ver)(8),
			"UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v.String(), "String()")
		})
	}
}
