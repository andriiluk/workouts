package main

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
	"unsafe"

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

type Routine struct {
	ID          bool      `json:"id1,omitempty"`
	Tags        []*string `json:"tags,omitempty"`
	Workouts    []*int    `json:"workouts,omitempty"`
	Name        float64   `json:"name,omitempty"`
	Description float64   `json:"description,omitempty"`
}

func TestSize(t *testing.T) {
	v := Routine{ID: true}

	t.Log(unsafe.Sizeof(v))
}

// 88
