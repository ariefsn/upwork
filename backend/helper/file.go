package helper

import (
	"errors"
	"os"
	"path"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ariefsn/upwork/constant"
)

func UploadGraphqlFile(file graphql.Upload, dir, alias string) (fileName string, filePath string, err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", "", err
		}
	}

	suffix := time.Now().Format("2006-01-02_15-04-05") + "_" + RandomString(6)

	contentType := strings.Split(file.ContentType, "/")
	ext := contentType[len(contentType)-1]
	fileName = file.Filename
	if alias != "" {
		fileName = alias
	}
	if len(strings.Split(fileName, ".")) > 1 {
		fileName = strings.Split(fileName, ".")[0]
	}
	fileName = fileName + "_" + suffix + "." + ext
	fileName = strings.ReplaceAll(fileName, " ", "_")
	filePath = path.Join(dir, fileName) // prepare file path

	buff := make([]byte, file.Size)
	_, err = file.File.Read(buff)
	if err != nil {
		return "", "", err
	}

	err = os.WriteFile(filePath, buff, 0644) // write buffer to file
	if err != nil {
		return "", "", err
	}

	isContainsUploads := strings.Contains(filePath, constant.DIR_UPLOAD)
	if isContainsUploads {
		filePath = strings.Replace(filePath, constant.DIR_UPLOAD, "", 1)
	}

	return fileName, filePath, nil
}

func IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveFile(filePath string) error {
	if !IsFileExists(filePath) {
		return errors.New("file not found")
	}

	return os.Remove(filePath)
}
