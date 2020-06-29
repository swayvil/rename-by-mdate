package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	timeFormat       = "2006-01-02_15.04.05"
	pathSeparatorStr = string(os.PathSeparator)
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("rename-by-mdate takes the root directory path in parameter")
	} else {
		dirName := ""
		dirPath := os.Args[1]
		currentTimeStamp := strconv.FormatInt(time.Now().UnixNano(), 10)

		if len(os.Args) == 3 {
			dirName = os.Args[2]
		}
		renameFiles(dirPath, dirName, currentTimeStamp)
	}
}

func renameFiles(dirPath string, dirName string, currentTimeStamp string) {
	file, err := os.Open(dirPath)
	if err != nil {
		log.Fatalf("Failed opening directory: %s", err)
	}
	defer file.Close()

	// Get the directory name
	if dirName == "" {
		i := strings.LastIndex(dirPath, pathSeparatorStr)
		dirName = dirPath[i+1 : len(dirPath)]
	}

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	n := 0
	for _, filename := range list {
		renameFile(dirPath+pathSeparatorStr, filename, dirName+"_", currentTimeStamp+"-"+strconv.Itoa(n))
		n++
	}
}

func renameFile(dirPath string, filename string, prefix string, tmpFileName string) {
	filePath := dirPath + filename

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if fi.IsDir() {
		return
	}

	var st syscall.Stat_t
	if err := syscall.Stat(filePath, &st); err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(st.Mtimespec.Sec, 0)
	newFilename := prefix + tm.Format(timeFormat) + strings.ToLower(filepath.Ext(filePath))

	os.Rename(filePath, dirPath+tmpFileName) // Rename to a temporary filename, so we are sure nameIfFileExists won't match with the file if it is already well named

	newFilePath := nameIfFileExists(dirPath + newFilename)
	os.Rename(dirPath+tmpFileName, newFilePath)
	fmt.Println(newFilePath)
}

// Check if file already exists, else increment a suffix number
func nameIfFileExists(filePath string) string {
	nb := 1
	newFilePath := filePath
	_, err := os.Stat(newFilePath)

	for !os.IsNotExist(err) {
		i := strings.LastIndex(filePath, ".")
		newFilePath = filePath[0:i] + "_" + strconv.Itoa(nb) + filePath[i:len(filePath)]
		_, err = os.Stat(newFilePath)
		nb++
	}
	return newFilePath
}
