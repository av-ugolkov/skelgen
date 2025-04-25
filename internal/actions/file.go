package actions

import (
	"fmt"
	"os"
	"path"
)

func CreateFile(rootPath, file string, value string) error {
	f, err := os.Create(path.Join(rootPath, file))
	if err != nil {
		return fmt.Errorf("couldn't create a file [%s] in the folder [%s]: %v", file, rootPath, err)
	}
	_, err = f.WriteString(value)
	if err != nil {
		return fmt.Errorf("couldn't write the data in the file [%s/%s]: %v", rootPath, file, err)
	}

	return nil
}
