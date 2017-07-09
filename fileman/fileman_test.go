package fileman

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFm_Search_WithNotExistsPath_ReturnsError(t *testing.T) {
	fm := NewFileMan()
	_, err := fm.Search("path/not/exists")
	assert.EqualError(t, err, "lstat path/not/exists: no such file or directory")
}

func TestFm_Search_WithExistsPath_ReturnsOneFile(t *testing.T) {
	fm := NewFileMan()
	actual, err := fm.Search("../tests/fixtures")
	assert.NoError(t, err)

	assert.Equal(t, 1, len(actual))
}

func Test_FmCalculate_WithFile_ReturnsSumFile(t *testing.T) {
	f, err := os.Open("../tests/fixtures/testfile.txt")
	if err != nil {
		t.Errorf("Cound not open file, %s", err)
	}

	fm := NewFileMan()
	actual, err := fm.Calculate(f)

	assert.NoError(t, err)
	assert.Equal(t, "69164f5d00c8ee3a8f6d67902689de94", actual)
}
