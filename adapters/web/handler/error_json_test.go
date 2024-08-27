package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "Error message"
	result := jsonError(msg)
	expected := `{"message":"Error message"}`
	require.Equal(t, []byte(expected), result)
}