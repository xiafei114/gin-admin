package export

import (
	"fmt"
	"gin-admin/internal/app/ginadmin/config"
)

const EXT = ".xlsx"

// GetExcelFullUrl get the full access path of the Excel file
func GetExcelFullUrl(hostName string, filePath string) string {
	return fmt.Sprintf("%s/%s", hostName, filePath)
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	cfg := config.GetGlobalConfig()
	return cfg.File.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	cfg := config.GetGlobalConfig()
	return cfg.File.RuntimeRootPath + GetExcelPath()
}
