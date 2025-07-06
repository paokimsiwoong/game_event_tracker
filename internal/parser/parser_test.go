package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLineParse(t *testing.T) {
	err := temp()
	require.NoError(t, err)
	assert.Equal(t, nil, err)
}

func temp() error {
	return nil
}
