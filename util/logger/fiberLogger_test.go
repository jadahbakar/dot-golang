package logger

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFiberLogger(t *testing.T) {
	// negative case
	file, err := NewFiberLogger("")
	assert.Error(t, err)
	assert.Nil(t, file)

	// positive case (bunting 25jt)
	currPath, _ := os.Getwd()
	logPath := filepath.Join(currPath, "../", "../", "log")
	file, err = NewFiberLogger(logPath)
	assert.NoError(t, err)
	assert.NotNil(t, file)
}
