package configit_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/michaelwomack/configit"
)

func TestLoad(t *testing.T) {
	type config struct {
		Float32  float32 `env:"FLOAT_32"`
		Float64  float64 `env:"FLOAT_64"`
		Int      int     `env:"INT"`
		Int8     int8    `env:"INT_8"`
		Int16    int16   `env:"INT_16"`
		Int32    int32   `env:"INT_32"`
		Int64    int64   `env:"INT_64"`
		Unsigned struct {
			Uint   uint   `env:"UINT"`
			Uint8  uint8  `env:"UINT_8"`
			Uint16 uint16 `env:"UINT_16"`
			Uint32 uint32 `env:"UINT_32"`
			Uint64 uint64 `env:"UINT_64"`
		}
		String string `env:"STRING"`
		Bool   bool   `env:"BOOL"`
	}
	envs := map[string]string{
		"FLOAT_32": "1.1",
		"FLOAT_64": "1.1",
		"INT":      "1",
		"INT_8":    "1",
		"INT_16":   "1",
		"INT_32":   "1",
		"INT_64":   "1",
		"UINT":     "1",
		"UINT_8":   "1",
		"UINT_16":  "1",
		"UINT_32":  "1",
		"UINT_64":  "1",
		"STRING":   "Test value",
		"BOOL":     "TRUE",
	}
	for key, value := range envs {
		require.NoError(t, os.Setenv(key, value))
		require.Equal(t, os.Getenv(key), value)
	}

	target := &config{}
	require.NoError(t, configit.Load(target))

	require.Equal(t, target.Float32, float32(1.1), "float32")
	require.Equal(t, target.Float64, float64(1.1), "float64")
	require.Equal(t, target.Int8, int8(1), "int8")
	require.Equal(t, target.Int16, int16(1), "int16")
	require.Equal(t, target.Int32, int32(1), "int32")
	require.Equal(t, target.Int64, int64(1), "int64")
	require.Equal(t, target.Unsigned.Uint, uint(1), "uint")
	require.Equal(t, target.Unsigned.Uint8, uint8(1), "uint8")
	require.Equal(t, target.Unsigned.Uint16, uint16(1), "uint16")
	require.Equal(t, target.Unsigned.Uint32, uint32(1), "uint32")
	require.Equal(t, target.Unsigned.Uint64, uint64(1), "uint64")
	require.Equal(t, target.String, "Test value", "string")
	require.Equal(t, target.Bool, true, "bool")
}