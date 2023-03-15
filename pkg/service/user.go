package service

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/anaxaim/tui/pkg/model"
	"github.com/anaxaim/tui/pkg/repository"
)

const (
	MinPasswordLength = 6
)

var ErrUserAlreadyExists = errors.New("user has already exist")

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) List() (model.Users, error) {
	return u.userRepository.List()
}

func (u *userService) Create(user *model.User) (*model.User, error) {
	exists, err := u.userRepository.Exists(user.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.Wrap(ErrUserAlreadyExists, fmt.Sprintf("username '%s'", user.Username))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)

	return u.userRepository.Create(user)
}

func (u *userService) Get(username string) (*model.User, error) {
	return u.getUserByUsername(username)
}

func (u *userService) Update(username string, newUser *model.User) (*model.User, error) {
	old, err := u.getUserByUsername(username)
	if err != nil {
		return nil, err
	}

	newUser.Username = old.Username
	newUser.ID = old.ID

	if len(newUser.Password) > 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		newUser.Password = string(password)
	}

	return u.userRepository.Update(newUser)
}

func (u *userService) Delete(username string) error {
	user, err := u.getUserByUsername(username)
	if err != nil {
		return err
	}

	return u.userRepository.Delete(user)
}

func (u *userService) Validate(user *model.User) error {
	if user == nil {
		return errors.New("user is empty")
	}
	if user.Username == "" {
		return errors.New("user name is empty")
	}
	if len(user.Password) < MinPasswordLength {
		return fmt.Errorf("password length must great than %d", MinPasswordLength)
	}
	return nil
}

func (u *userService) getUserByUsername(username string) (*model.User, error) {
	return u.userRepository.GetUserByUsername(username)
}
