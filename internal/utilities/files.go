package utilities

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		parentDir := filepath.Dir(path)
		info, err := os.Stat(parentDir)
		if err != nil {
			return fmt.Errorf("could not get info of parent directory: %v", err)
		}
		return os.MkdirAll(path, info.Mode())
	} else if err != nil {
		return err
	}
	return nil
}
