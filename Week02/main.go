package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	xerrors "github.com/pkg/errors"
)

/**
 * File :   main.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/11/29 21:38
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

var Engine *gorm.DB

type User struct {
	ID int64
}

func (User) TableName() string {
	return "biz_user"
}

func (u *User) NoDataError() error {
	return sql.ErrNoRows
}

func (u *User) GetUser(id int64) (*User, error) {
	var doc *User

	table := Engine.Table(u.TableName())
	if u.ID != 0 {
		table = table.Where("id = ?", id)
	}

	// 应该将 not found 翻译为明确指令给到上层，因为对于 DAO 层，不同 DB 的 lib 库可能有区别
	if err := table.First(doc).Error; err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return doc, xerrors.WithMessage(err, "no data found")
		}
		return nil, xerrors.Wrap(err, "db error")
	}
	return doc, nil
}

func main() {
	var u User
	user, err := u.GetUser(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("user info: %+v\n", user)
}
