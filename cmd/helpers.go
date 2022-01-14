package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
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
