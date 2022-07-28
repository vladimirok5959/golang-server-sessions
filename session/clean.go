package session

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Keep session files for 7 days
const LifeTime = (24 * time.Hour) * 7

// Clean to remove all local session files which not was used more than 24 hours
func Clean(tmpdir string) error {
	files, err := ioutil.ReadDir(tmpdir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if len(file.Name()) == 40 {
			if diff := time.Since(file.ModTime()); diff > LifeTime {
				if err := os.Remove(strings.Join(
					[]string{tmpdir, file.Name()},
					string(os.PathSeparator),
				)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
