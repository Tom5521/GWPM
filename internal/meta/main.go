package meta

import (
	"strconv"
)

var (
	Version     string
	BuildNumber string
	DevBuild    = func() bool {
		b, _ := strconv.ParseBool(DevBuildStr)
		return b
	}()
	Release = func() bool {
		b, _ := strconv.ParseBool(ReleaseStr)
		return b
	}()
)

// These variables represent a Boolean value that will be assigned when compiling
var (
	DevBuildStr string
	ReleaseStr  string
)
