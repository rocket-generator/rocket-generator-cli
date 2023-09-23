package a_download_template

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var templates = map[string]string{
	"go-gin":      "https://github.com/rocket-generator/go-gin-template/archive/refs/heads/master.zip",
	"php-laravel": "https://github.com/rocket-generator/php-laravel-template/archive/refs/heads/master.zip",
	"react-admin": "https://github.com/rocket-generator/typescript-react-admin-template/archive/refs/heads/master.zip",
}

func getDownloadURL(templateName string) (*string, error) {
	lowerTemplateName := strings.ToLower(templateName)
	url, ok := templates[lowerTemplateName]
	if !ok {
		if _, err := os.Stat(templateName); err != nil {
			return nil, errors.New("template not found")
		}
		absolutePath, err := filepath.Abs(templateName)
		if err != nil {
			return nil, err
		}
		return &absolutePath, nil
	}

	return &url, nil
}

func downloadFile(targetURL string, directoryPath string) (string, error) {
	response, err := grab.Get(directoryPath, targetURL)
	if err != nil && response.HTTPResponse.StatusCode != 200 {
		log.Fatal(err)
		return "", err
	}
	return response.Filename, nil
}

// Source: https://github.com/artdarek/go-unzip/blob/master/unzip.go
func unzip(zipFilePath string, destinationFilePath string) error {

	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	err = os.MkdirAll(destinationFilePath, 0755)
	if err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(destinationFilePath, f.Name)
		if !strings.HasPrefix(path, filepath.Clean(destinationFilePath)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: Illegal file path", path)
		}

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(path, f.Mode())
		} else {
			_ = os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
