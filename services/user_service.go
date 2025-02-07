// services/user_service.go
package services

import (
    "errors"

    "github.com/chatx/models"
    "github.com/chatx/repositories"
)

type UserService struct {
    userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

func (us *UserService) GetUser(id string) (*models.User, error) {
    return us.userRepository.GetUser(id)
}

