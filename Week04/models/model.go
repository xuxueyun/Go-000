package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt" gorm:"index;comment:'创建时间'"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"comment:'更新时间'"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index;comment:'删除时间'"`
}
