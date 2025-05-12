package regex

import "fmt"

var (
	ErrInvalidFileName   = fmt.Errorf("invalid file name")
	ErrInvalidFolderName = fmt.Errorf("invalid folder name")
)

func IsValidate(s string) bool {
	return fileFolderName.MatchString(s)
}
