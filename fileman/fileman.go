package fileman

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileMan is the interface that wraps the basic Search and Calculate methods
type FileMan interface {
	Search(path string) (files map[string]*os.File, err error)
	Calculate(f *os.File) (sum string, err error)
}

type fm struct {
}

func NewFileMan() FileMan {
	return &fm{}
}

func (fm *fm) Search(root string) (files map[string]*os.File, err error) {
	files = map[string]*os.File{}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		files[filepath.Base(f.Name())] = f

		return nil
	})

	return
}

func (fm *fm) Calculate(f *os.File) (sum string, err error) {
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	sum = fmt.Sprintf("%x", hash.Sum(nil))

	return
}
