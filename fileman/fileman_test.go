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

	assert.Equal(t, 44, len(actual))
}

func Test_FmCalculate_WithFile_ReturnsSumFile(t *testing.T) {
	f, err := os.Open("../tests/fixtures/out_1_44.sql")
	if err != nil {
		t.Errorf("Cound not open file, %s", err)
	}

	fm := NewFileMan()
	actual, err := fm.Calculate(f)

	assert.NoError(t, err)
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", actual)
}
