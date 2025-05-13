package actions

import (
	"fmt"
	"os"
	"path"
	"strings"

	keywords "github.com/av-ugolkov/yask/internal/key-words"
	"github.com/av-ugolkov/yask/internal/regex"
)

func CreateFolders(rootPath, folder string) (string, error) {
	folder = keywords.RemoveInsulator(folder)

	var folders []string
	if HasSubfolders(folder) {
		folders = strings.Split(folder, "/")
	} else {
		folders = []string{folder}
	}

	for _, subFolder := range folders {
		err := createFolder(rootPath, subFolder)
		if err != nil {
			return "", err
		}
		rootPath = path.Join(rootPath, subFolder)
	}

	return rootPath, nil
}

func createFolder(rootPath, folder string) error {
	if !regex.IsValidate(folder) {
		return fmt.Errorf("%v: [%s]", regex.ErrInvalidFolderName, folder)
	}
	err := os.Mkdir(path.Join(rootPath, folder), 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("couldn't create the folder [%s] by the path [%s]: %v", folder, rootPath, err)
	}
	return nil
}

func HasSubfolders(folder string) bool {
	return strings.Contains(folder, "/")
}
