package perm

import "os"

var IsAdmin bool = func() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}()
