package version

import (
	"encoding/json"
	"runtime"
)

var (
	gitVersion = "v0.0.0-master+$Format:%H$"

	buildDate = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')

	version = &Version{
		Version:   gitVersion,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
	}
)

type Version struct {
	Version   string `json:"version"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
}

func (v *Version) String() string {
	data, _ := json.Marshal(v) //nolint: errchkjson
	return string(data)
}

func Get() *Version {
	return version
}
