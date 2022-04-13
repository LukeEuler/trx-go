package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyFromHex(t *testing.T) {
	tests := []struct {
		name        string
		hexKey      string
		wantAddress string
		wantErr     bool
	}{
		{
			"test 1",
			"11ced275270a23278b76c756ebe566650db48dd854b905c15cce8e7c123c73c0",
			"TLGf8u9G38EnF7HBQuWfCpnQKXW5NbXEV1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKeyFromHex(tt.hexKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyFromHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.hexKey, got.PrivateKey())
			assert.Equal(t, tt.wantAddress, got.Address())
		})
	}
}
