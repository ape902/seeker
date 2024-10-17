package versionx

import "fmt"

var (
	Version   string
	BuildTime string
)

func GetVersion() string {
	return fmt.Sprintf(": %s (build time: %s)", Version, BuildTime)
}
