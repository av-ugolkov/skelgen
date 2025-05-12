//go:build windows

package regex

import "regexp"

var fileFolderName = regexp.MustCompile(`^(?i)(?!^(CON|PRN|AUX|NUL|COM[1-9]|LPT[1-9])(\..*)?$)[^<>:"/\\|?*\x00-\x1F]+[^. ]$`)
