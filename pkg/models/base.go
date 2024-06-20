package models

import "github.com/dgrijalva/jwt-go"

type (
	CustomClaims struct {
		ID       uint
		NickName string
		jwt.StandardClaims
	}

	BaseModel struct {
		Id        int   `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
		CreatedAt int64 `gorm:"column:created_at" json:"created_at" form:"created_at"`
		UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
		DeletedAt int64 `gorm:"column:deleted_at" sql:"index" json:"-"`
	}

	PageModel struct {
		Size  int64 `json:"size"`
		Index int64 `json:"index"`
	}

	IDS struct {
		IDS []int `json:"ids"`
	}
)
