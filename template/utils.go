package template

import (
	"path/filepath"

	"github.com/coveo/gotemplate/v3/utils"
)

func getTargetFile(targetFile, sourcePath, targetPath string) string {
	if targetPath != "" {
		targetFile = filepath.Join(targetPath, utils.Relative(sourcePath, targetFile))
	}
	return targetFile
}
