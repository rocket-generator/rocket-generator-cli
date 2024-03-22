package utilities

import (
	"github.com/stoewer/go-strcase"
	"path/filepath"
	"strings"
)

func GetFilenameWithoutExtension(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	return name
}

func RemovePostfix(s string, postfix string) string {
	lower := strcase.LowerCamelCase(s)
	if strings.HasSuffix(lower, postfix) {
		return strings.TrimSuffix(s, postfix)
	}
	return s
}
