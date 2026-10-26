package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
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

func (s *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	//cek jika user id tidak sesuai
	if user.ID == 0 {
		return user, errors.New("user not found on that email")
	}
	//compare passwordnya
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil //agar email tidak double jika exists
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {

	//ambil user by ID
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	//update attribute avatar filename
	user.AvatarFileName = fileLocation

	//simpan perubahan avatat filename
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
