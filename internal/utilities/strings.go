package utilities

import "path/filepath"

func GetFilenameWithoutExtension(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(ext)]
	return name
}
