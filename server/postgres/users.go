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

// Batch function for DataLoader - fetches users by multiple IDs
func (m *UsersRepo) GetUsersByIDs(ids []string) ([]*model.User, []error) {
	users := make([]*model.User, len(ids))
	errs := make([]error, len(ids))

	var dbUsers []*model.User
	err := m.DB.Model(&dbUsers).Where("id IN (?)", pg.In(ids)).Select()
	if err != nil {
		// Return error for all IDs
		for i := range errs {
			errs[i] = err
		}
		return users, errs
	}

	// Map users to their positions based on requested IDs
	userMap := make(map[string]*model.User)
	for _, user := range dbUsers {
		userMap[user.ID] = user
	}

	for i, id := range ids {
		users[i] = userMap[id]
	}

	return users, errs
}

func (m *UsersRepo) CreateUser(input model.NewUserInput) (*model.User, error) {
	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName, // Fixed: was input.FirstName
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
