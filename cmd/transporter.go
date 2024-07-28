package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	config "github.com/freeflowuniverse/gitea-rebrander/internal"
)

type Transporter struct {
	config *config.Config
}

func NewTransporter(config *config.Config) *Transporter {
	return &Transporter{config: config}
}

func (t *Transporter) Transport() error {
	root := t.config.GiteaRepoPath

	for fileName, targetFilePath := range t.config.Files {
		targetFilePath = t.getTargetFilePath(targetFilePath)
		sourceFilePath := t.getSourceFilePath(fileName)

		sourceContent, err := os.ReadFile(sourceFilePath)
		if err != nil {
			return err
		}

		err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && path == targetFilePath {
				log.Printf("found target file %s\n", path)

				err = os.WriteFile(targetFilePath, sourceContent, 0644)
				if err != nil {
					return err
				}

				log.Println("target file content overridden successfully.")
				return filepath.SkipDir
			}

			return nil
		})

		if err != nil {
			log.Fatalf("error walking the directory %v", err)
		}
	}
	return nil
}

func (t *Transporter) getTargetFilePath(targetFilePath string) string {
	return fmt.Sprintf("%s/%s", t.config.GiteaRepoPath, targetFilePath)
}

func (t *Transporter) getSourceFilePath(fileName string) string {
	return fmt.Sprintf("./files/%s", fileName)
}
