package generator

import (
	"errors"
	"sync"

	"github.com/av-ugolkov/gopkg/logger"
	"github.com/av-ugolkov/gopkg/safe"
	"github.com/av-ugolkov/yask/internal/actions"
	"github.com/av-ugolkov/yask/internal/config"
	kw "github.com/av-ugolkov/yask/internal/key-words"
)

func GenSkeleton(pathFile string, instance any, dynamic map[string]string) error {
	err := config.Load(pathFile, dynamic)
	if err != nil {
		return err
	}

	var listOfErrors error
	chErr := startGenerate(config.Skel(), ".")
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
		case string(kw.Exec):
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
		case string(kw.Dirs):
			for _, d := range v.([]any) {
				_, err := actions.CreateFolders(rootPath, d.(string))
				if err != nil {
					chErr <- err
				}
			}
		case string(kw.Files):
			data := v.([]any)
			for _, d := range data {
				switch d.(type) {
				case map[string]any:
					wg.Add(1)
					safe.Go(func() {
						defer wg.Done()
						err := <-startGenerate(d.(map[string]any), rootPath)
						if err != nil {
							chErr <- err
						}
					})
				case string:
					err := actions.CreateFile(rootPath, d.(string), "")
					if err != nil {
						chErr <- err
					}
				}
			}
		default:
			switch v.(type) {
			case map[string]any:
				createdPath, err := actions.CreateFolders(rootPath, k)
				if err != nil {
					chErr <- err
					continue
				}
				wg.Add(1)
				safe.Go(func() {
					defer wg.Done()
					err = <-startGenerate(v.(map[string]any), createdPath)
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
