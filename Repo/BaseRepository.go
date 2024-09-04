package Repo

import (
	"ESayur/Model"
	"context"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	ConnDB *gorm.DB
}

func NewMysqlRepo(connDB *gorm.DB) MysqlRepo {
	return &mysqlRepo{connDB}
}

type MysqlRepo interface {
	// User
	GetUsers(ctx context.Context) (users []*Model.User, err error)
}
