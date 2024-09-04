package Services

import (
	"ESayur/Model"
	"ESayur/Repo"
	"context"
	"errors"
)

type userServices struct {
	mysqlRepo Repo.MysqlRepo
}

func NewUserService(mysqlRepo Repo.MysqlRepo) UserServices {
	return &userServices{mysqlRepo}
}

type UserServices interface {
	GetUsers(ctx context.Context) (users []*Model.UserResponse, err error)
}

func (us *userServices) GetUsers(ctx context.Context) (users []*Model.UserResponse, err error) {
	data, err := us.mysqlRepo.GetUsers(ctx)

	if len(data) == 0 {
		return nil, errors.New("data kosong")
	}

	for _, user := range data {
		users = append(users, &Model.UserResponse{
			UserName: user.Username,
			Role:     user.RoleUser.NamaRole,
			Cart:     user.Cart,
			Wallet:   user.Wallet,
		})
	}
	return
}
