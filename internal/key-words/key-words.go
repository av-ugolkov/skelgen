package keywords

import "strings"

type KeyWords string

const (
	Skel      KeyWords = "skel"
	Content   KeyWords = "content"
	Dirs      KeyWords = "dirs"
	Files     KeyWords = "files"
	Link      KeyWords = "link"
	Exec      KeyWords = "exec"
	Insulator KeyWords = "^"
)

func RemoveInsulator(s string) string {
	return strings.TrimPrefix(s, string(Insulator))
}
