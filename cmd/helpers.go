package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var PathSeparatorString = string(os.PathSeparator)

// MD5Dir reads all the files in the file tree rooted at root and return an md5
// of each file name and a serialized appended string of its contents
func MD5Dir(root string) (string, error) {
	var dataConcatenation []byte
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		dataConcatenation = []byte(string(dataConcatenation) + path + string(data))

		return nil
	})

	if err != nil {
		return "", err
	}

	hash := md5.Sum(dataConcatenation)
	return hex.EncodeToString(hash[:]), nil
}

// MD5File reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5File(root string) (string, error) {
	content, err := ioutil.ReadFile(root)
	if err != nil {
		log.Fatal(err)
	}
	hash := md5.Sum(content)

	// Convert []byte to string and print to screen
	return hex.EncodeToString(hash[:]), nil
}

func IsHydroRepo(path string) bool {
	for _, dir := range initialDirs {
		_, err := os.Stat(path + PathSeparatorString + dir)
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func InHydroRepo(path string) bool {
	if path == filepath.Base(".") || !fs.ValidPath(path) {
		return false
	}
	if IsHydroRepo(path) {
		return true
	}
	return InHydroRepo(filepath.Dir(path))
}
