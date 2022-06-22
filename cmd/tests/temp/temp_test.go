package main

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeError(t *testing.T) {
	type ErrStr struct {
		Err error `json:"err,omitempty"`
	}

	var errSt = ErrStr{errors.New("test error")}

	err := json.NewEncoder(os.Stdout).Encode(errSt)
	require.NoError(t, err)
}
