package service

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/repository"
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
	exists, err := u.userRepository.Exists(user.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.Wrap(ErrUserAlreadyExists, fmt.Sprintf("name '%s'", user.Name))
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(password)

	return u.userRepository.Create(user)
}

func (u *userService) Get(name string) (*model.User, error) {
	return u.getUserByName(name)
}

func (u *userService) Update(name string, newUser *model.User) (*model.User, error) {
	old, err := u.getUserByName(name)
	if err != nil {
		return nil, err
	}

	newUser.Name = old.Name
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

func (u *userService) Delete(name string) error {
	user, err := u.getUserByName(name)
	if err != nil {
		return err
	}

	return u.userRepository.Delete(user)
}

func (u *userService) Validate(user *model.User) error {
	if user == nil {
		return errors.New("user is empty")
	}
	if user.Name == "" {
		return errors.New("user name is empty")
	}
	if len(user.Password) < MinPasswordLength {
		return fmt.Errorf("password length must great than %d", MinPasswordLength)
	}
	return nil
}

func (u *userService) Auth(auser *model.AuthUser) (*model.User, error) {
	if auser == nil || auser.Name == "" || auser.Password == "" {
		return nil, fmt.Errorf("name or password is empty")
	}

	user, err := u.userRepository.GetUserByName(auser.Name)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auser.Password)); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (u *userService) getUserByName(name string) (*model.User, error) {
	return u.userRepository.GetUserByName(name)
}
