package internal

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/av-ugolkov/gopkg/logger"
	"gopkg.in/yaml.v3"
)

func GenSkeleton(pathFile string, instance any) error {
	f, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var mapConfig map[string]any
	err = yaml.Unmarshal(f, &mapConfig)
	if err != nil {
		return err
	}

	return startGenerate(mapConfig, ".")
}

func startGenerate(conf map[string]any, rootPath string) error {
	var listErrors error

	for k, v := range conf {
		switch k {
		case string(Dirs):
			logger.Infof("create dirs: not implemented yet")
		case string(Files):
			logger.Infof("create files: not implemented yet")
		case string(Exec):
			commands := v.([]any)
			for _, command := range commands {
				cmd := strings.Split(command.(string), " ")
				runInDir(rootPath, cmd[0], cmd[1:]...)
			}
		case string(Link):
			logger.Infof("link file: not implemented yet")
		default:
			switch v.(type) {
			case map[string]any:
				err := createFolder(rootPath, k)
				if err != nil {
					listErrors = errors.Join(listErrors, err)
				}
				err = startGenerate(v.(map[string]any), path.Join(rootPath, k))
				if err != nil {
					listErrors = errors.Join(listErrors, err)
				}
			case string:
				err := createFile(rootPath, k, v.(string))
				if err != nil {
					listErrors = errors.Join(listErrors, err)
				}
			default:
				logger.Warnf("unknow command - [%s]: %v", k, v)
			}
		}
	}

	return listErrors
}
