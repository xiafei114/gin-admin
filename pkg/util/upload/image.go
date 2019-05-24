package upload

import (
	"fmt"
	"go-gin-example/pkg/logging"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"gin-admin/internal/app/ginadmin/config"
	"gin-admin/pkg/util"
	"gin-admin/pkg/util/file"
)

// GetFileFullUrl 文件下载url
func GetFileFullUrl(hostName string, filePath string) string {
	return fmt.Sprintf("%s/%s", hostName, filePath)
}

// GetFileName get file name
func GetFileName(name string) (string, string) {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	// fileName = util.EncodeMD5(fileName)
	fileName = util.MustUUID()

	return fileName, ext
}

// GetFilePath 返回全路径和相对地址
func GetFilePath(alias string) (string, string) {
	cfg := config.GetGlobalConfig()
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	relativePath := fmt.Sprintf("%s%s%s/", cfg.File.ImageSavePath, alias, tm.Format("2006-01-02"))
	return fmt.Sprintf("%s%s", cfg.File.RuntimeRootPath, relativePath), relativePath
}

// GetFileFullPath 获得全路径
func GetFileFullPath(relative string) string {
	cfg := config.GetGlobalConfig()
	return fmt.Sprintf("%s%s", cfg.File.RuntimeRootPath, relative)
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	cfg := config.GetGlobalConfig()
	for _, allowExt := range cfg.File.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckFileSize check file size
func CheckFileSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	cfg := config.GetGlobalConfig()

	return size <= cfg.File.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
