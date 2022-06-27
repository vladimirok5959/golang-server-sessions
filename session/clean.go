package session

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Clean(tmpdir string) error {
	files, err := ioutil.ReadDir(tmpdir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if len(file.Name()) == 40 {
			if diff := time.Since(file.ModTime()); diff > 24*time.Hour {
				err = os.Remove(strings.Join([]string{tmpdir, file.Name()}, string(os.PathSeparator)))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
