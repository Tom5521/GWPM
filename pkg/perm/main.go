package perm

import "os"

// A module only for this?
var IsAdmin bool = func() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}()
