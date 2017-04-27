// +build !windows

package lib

import "os"

func IsFilePermissionTooOpen(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if s.Mode()&077 != 0 {
		return true, nil
	}

	return false, nil
}