package internal

import (
	"fmt"
	"os"
	"path"
)

func createFolder(rootPath, folder string) error {
	err := os.Mkdir(path.Join(rootPath, folder), 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("Folder [%s] wasn't created: %v", folder, err)
	}
	return nil
}
