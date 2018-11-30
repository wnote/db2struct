package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func GetPackageNameFromOutPutDir(outPutDir string) string {
	outPutDir, err := filepath.Abs(outPutDir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	packageName := filepath.Base(outPutDir)
	packageName = regexp.MustCompile("^[^a-zA-Z]+").ReplaceAllString(packageName, "")
	packageName = regexp.MustCompile("[^a-zA-Z0-9_]+").ReplaceAllString(packageName, "")
	return packageName
}

func CheckAndMakeDir(dir string, fileMode string) string {
	exist, err := pathExists(dir)
	if err != nil {
		panic(err)
	}
	if !exist {
		err := os.MkdirAll(dir, getFileModeFromString(fileMode))
		if err != nil {
			fmt.Printf("%s", err)
			panic(err)
		} else {
			fmt.Print("Make dir OK!")
		}
	}
	return dir
}

func CreateAndWriteFile(fileName string, content string, fileMode string) {
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(fmt.Sprintf("Create file %s err:%v", fileName, err))
		panic(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	fp.Chmod(getFileModeFromString(fileMode))
	fp.WriteString(content)
	fp.Close()
	fmt.Println("Create model file " + filepath.Base(fileName))
}

func getFileModeFromString(fileMode string) os.FileMode {
	fileModeNu, err := strconv.ParseInt(fileMode, 8, 16)
	if err != nil {
		fmt.Println(fmt.Sprintf("fileMode %s err:%v", fileMode, err))
		panic(err)
	}
	return os.FileMode(fileModeNu)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileContent(filePath string) []byte {
	exist, err := pathExists(filePath)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic(filePath + " not exist!")
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return fileContent
}
