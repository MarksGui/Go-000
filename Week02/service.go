package main

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (p *UserService) GetInfo() (*UserEntity, error) {
	return NewUserRepository().GetInfo()
}
