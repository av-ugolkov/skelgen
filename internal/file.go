package internal

import (
	"os"
	"path"
)

func createFile(rootPath, file string, value string) error {
	f, err := os.Create(path.Join(rootPath, file))
	if err != nil {
		return err
	}
	_, err = f.WriteString(value)
	if err != nil {
		return err
	}

	return nil
}
