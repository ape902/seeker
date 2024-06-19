package cmdb

import "github.com/ape902/seeker/pkg/models"

type (
	HostInfo struct {
		models.BaseModel
		IP    string `json:"ip"`
		Port  int    `json:"port"`
		OS    string `json:"os"`
		Label string `json:"label"`
	}
)
