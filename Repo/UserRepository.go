package Repo

import (
	"ESayur/Model"
	"context"
)

func (sql *mysqlRepo) GetUsers(ctx context.Context) (users []*Model.User, err error) {
	err = sql.ConnDB.WithContext(ctx).Find(&users).Error

	return
}
