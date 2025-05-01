package lib

import (
	"path"
	"strings"
)

func FileName(p string, e string) string {
	name := path.Base(p)
	ext := path.Ext(name)
	baseName := strings.TrimSuffix(name, ext)

	return baseName + e
}
