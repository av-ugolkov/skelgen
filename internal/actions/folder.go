package actions

import (
	"fmt"
	"os"
	"path"
)

func CreateFolder(rootPath, folder string) error {
	err := os.Mkdir(path.Join(rootPath, folder), 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("couldn't create the folder [%s] by the path [%s]: %v", folder, rootPath, err)
	}
	return nil
}
