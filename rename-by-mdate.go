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
	if len(os.Args) <= 2 {
		log.Fatal("usage: rename-by-mdate <source directory path> <target directory path> [optional suffix]")
	} else {
		dirName := ""
		srcDirPath := os.Args[1]
		currentTimeStamp := strconv.FormatInt(time.Now().UnixNano(), 10)
		destDirPath := os.Args[2]

		if len(os.Args) == 4 {
			dirName = os.Args[3]
		}
		renameFiles(srcDirPath, dirName, currentTimeStamp, destDirPath)
	}
}

func renameFiles(srcDirPath string, dirName string, currentTimeStamp string, destDirPath string) {
	file, err := os.Open(srcDirPath)
	if err != nil {
		log.Fatalf("Failed opening directory: %s", err)
	}
	defer file.Close()

	// Get the directory name
	if dirName == "" {
		i := strings.LastIndex(srcDirPath, pathSeparatorStr)
		dirName = srcDirPath[i+1 : len(srcDirPath)]
	}

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	n := 0
	for _, filename := range list {
		renameFile(srcDirPath, destDirPath, filename, "_"+dirName, currentTimeStamp+"-"+strconv.Itoa(n))
		n++
	}
}

func renameFile(srcDirPath string, destDirPath string, filename string, suffix string, tmpFileName string) {
	filePath := filepath.Join(srcDirPath, filename)

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
	newFilename := tm.Format(timeFormat) + suffix + strings.ToLower(filepath.Ext(filePath))

	os.Rename(filePath, srcDirPath+tmpFileName) // Rename to a temporary filename, so we are sure nameIfFileExists won't match with the file if it is already well named

	newPathDir := filepath.Join(destDirPath, strconv.Itoa(tm.Year()))
	if err := os.MkdirAll(newPathDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	newFilePath := nameIfFileExists(filepath.Join(newPathDir, newFilename))
	os.Rename(srcDirPath+tmpFileName, newFilePath)
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
