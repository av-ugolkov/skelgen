package actions

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func CreateFolder(rootPath, folder string) error {
	err := os.Mkdir(path.Join(rootPath, folder), 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("couldn't create the folder [%s] by the path [%s]: %v", folder, rootPath, err)
	}
	return nil
}

func CreateFolders(rootPath string, folders ...string) error {
	for _, folder := range folders {
		if strings.Contains(folder, "/") {
			err := createSubfolders(rootPath, strings.Split(folder, "/")...)
			if err != nil && !os.IsExist(err) {
				return fmt.Errorf("couldn't create the folder [%s] by the path [%s]: %v", folder, rootPath, err)
			}
		} else {
			err := CreateFolder(rootPath, folder)
			if err != nil {
				return fmt.Errorf("couldn't create the folder [%s] by the path [%s]: %v", folder, rootPath, err)
			}
		}
	}
	return nil
}

func createSubfolders(rootPath string, folders ...string) error {
	if len(folders) == 0 {
		return nil
	}

	err := CreateFolder(rootPath, folders[0])
	if err != nil {
		return err
	}
	return createSubfolders(path.Join(rootPath, folders[0]), folders[1:]...)
}
