package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

// internal struct
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

/**

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:         input.Name,
		Email:        input.Email,
		Occupation:   input.Occupation,
		PasswordHash: string(passwordHash),
		Role:         "user",
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}


*/

// metode 1
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	//generate password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash) //convert ke string karena saat hash jd byte
	user.Role = "user"
	//save to database
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}
