package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/av-ugolkov/gopkg/logger"
)

func genFolders(projectMap map[string]any) error {
	for key, value := range projectMap {
		switch key {
		case "^exec":
			return runCmdInRootDir(value.(string))
		default:
			err := os.Mkdir(key, 0755)
			if err != nil && !os.IsExist(err) {
				return err
			}
			createSubfolders(key, value)
		}
	}

	return nil
}

func createSubfolders(parent string, value any) {
	if valueMap, ok := value.(map[string]any); ok {
		for key, subValue := range valueMap {
			switch key {
			case "^exec":
				cmds := subValue.([]any)
				for _, cmd := range cmds {
					c := strings.Split(cmd.(string), " ")
					err := runCmd(parent, c[0], c[1:]...)
					if err != nil {
						logger.Errorf("%v", err)
					}
				}
			default:
				if filepath.Ext(key) != "" {
					f, err := os.Create(path.Join(parent, key))
					if err != nil {
						logger.Errorf("%v", err)
					}
					_, err = f.WriteString(subValue.(string))
					if err != nil {
						logger.Errorf("%v", err)
					}
				} else {
					err := os.Mkdir(path.Join(parent, key), 0755)
					if err != nil && !os.IsExist(err) {
						logger.Errorf("%v", err)
					}

					createSubfolders(path.Join(parent, key), subValue)
				}
			}
		}
	}
}
