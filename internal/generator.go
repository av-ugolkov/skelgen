package internal

import (
	"errors"
	"os"
	"path"
	"sync"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/gopkg/safe"
	"github.com/av-ugolkov/yask/internal/actions"

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

	var listOfErrors error
	chErr := startGenerate(mapConfig, ".")
	for err := range chErr {
		listOfErrors = errors.Join(listOfErrors, err)
	}
	return listOfErrors
}

func startGenerate(conf map[string]any, rootPath string) chan error {
	var wg sync.WaitGroup
	chErr := make(chan error, 1)

	for k, v := range conf {
		switch k {
		case string(Dirs):
			logger.Infof("create dirs: not implemented yet")
		case string(Files):
			logger.Infof("create files: not implemented yet")
		case string(Exec):
			wg.Add(1)
			safe.Go(func() {
				defer wg.Done()

				commands := v.([]any)
				for _, command := range commands {
					err := actions.ExecCmdInDir(rootPath, command.(string))
					if err != nil {
						chErr <- err
						return
					}
				}

			})
		case string(Link):
			logger.Infof("link file: not implemented yet")
		default:
			switch v.(type) {
			case map[string]any:
				err := actions.CreateFolder(rootPath, k)
				if err != nil {
					chErr <- err
					continue
				}
				wg.Add(1)
				safe.Go(func() {
					defer wg.Done()

					err = <-startGenerate(v.(map[string]any), path.Join(rootPath, k))
					if err != nil {
						chErr <- err
					}
				})
			case string:
				err := actions.CreateFile(rootPath, k, v.(string))
				if err != nil {
					chErr <- err
				}
			default:
				logger.Warnf("unknow command - [%s]: %v", k, v)
			}
		}
	}
	safe.Go(func() {
		wg.Wait()
		close(chErr)
	})

	return chErr
}
