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
	Dynamic     KeyWords = "$"
)

func RemoveInsulator(s string) string {
	return strings.TrimPrefix(s, string(Insulator))
}

func IsDynamic(s string) bool {
	return strings.HasPrefix(s, string(Dynamic))
}
