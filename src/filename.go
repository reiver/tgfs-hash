package src

import (
	"github.com/reiver/tgfs-hash/cl"

	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	tempFileNamePattern = "tgfs-stdin-*.temp"
)

func FileName() (string, error) {

	{
		stat, _ := os.Stdin.Stat()
		if 0 == (stat.Mode() & os.ModeCharDevice) {
			tempFile, err := ioutil.TempFile("", tempFileNamePattern)
			if nil != err {
				err = fmt.Errorf("ERROR: Could not create temp file to store contents from STDIN: %s", err)
				return "", err
			}
			defer tempFile.Close()

			_, err = io.Copy(tempFile, os.Stdin)
			if nil != err {
				err = fmt.Errorf("ERROR: Could not store contents from STDIN in temp file: %s", err)
				return "", err
			}

			return tempFile.Name(), nil

		}
	}

	if 1 <= len(cl.Args) {

		return cl.Args[0], nil
	}

	return "", ErrFileNameNotFound
}
