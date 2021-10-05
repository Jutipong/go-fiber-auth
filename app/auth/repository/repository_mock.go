package repository

import (
	"auth/app/auth/model"
)

type repositoryMock struct {
	auth *model.Auth
}

func NewRepositoryMock(auth *model.Auth) IRepository {
	return &repositoryMock{auth}
}

func (r *repositoryMock) Inquiry_Auth(userName string) (result model.Auth, err error) {
	return *r.auth, nil
}

func (r *repositoryMock) Create_UserAndAuth(auth *model.Auth, user *model.User) error {
	return nil
}
