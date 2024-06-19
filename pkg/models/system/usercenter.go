package system

import "github.com/ape902/seeker/pkg/models"

type (
	User struct {
		models.BaseModel
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		NickName string `json:"nick_name"`
		Rule     int    `json:"rule"`
	}

	PasswordLoginFrom struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
)

func (User) TableName() string {
	return "system_user"
}
