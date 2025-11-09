package postgres

import (
	"github.com/SBPH-Matthew/testosterone-tracker/graph/model"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (m *UsersRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := m.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UsersRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UsersRepo) CreateUser(input model.NewUserInput) (*model.User, error) {
	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.FirstName,
		Gender:    input.Gender,
		Password:  input.Password,
		Email:     input.Email,
		Age:       input.Age,
	}

	_, err := m.DB.Model(user).Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}
