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

var (
	DevBuildStr string
	ReleaseStr  string
)
