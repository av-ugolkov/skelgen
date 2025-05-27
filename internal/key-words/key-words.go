package keywords

import (
	"strings"
)

type KeyWords string

const (
	Skel        KeyWords = "skel"
	Content     KeyWords = "content"
	Dirs        KeyWords = "dirs"
	Files       KeyWords = "files"
	ContentLink KeyWords = "#"
	FileLink    KeyWords = "@"
	Exec        KeyWords = "exec"
	Insulator   KeyWords = "^"
	Placeholder KeyWords = "$"
)

func RemoveInsulator(s string) string {
	return strings.TrimPrefix(s, string(Insulator))
}
