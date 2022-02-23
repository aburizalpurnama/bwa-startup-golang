package user

import "golang.org/x/crypto/bcrypt"

// create interface
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

// create implement struct
type service struct {
	repository Repository
}

// create instance method
func NewService(repository Repository) *service {
	return &service{repository: repository}
}

// add method to struct
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passHashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passHashed)

	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, err
}
