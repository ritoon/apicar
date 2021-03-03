package db

import (
	"github.com/google/uuid"

	"apicar/model"
)

var listUser map[string]*model.User = make(map[string]*model.User)

func GetUser(id string) (*model.User, error) {
	if len(id) == 0 {
		return nil, model.ErrAPIBadRequest
	}

	if u, ok := listUser[id]; ok {
		return u, nil
	}

	return nil, model.ErrAPINotfound
}

func CreateUser(u *model.User) (*model.User, error) {
	u.ID = uuid.New().String()
	listUser[u.ID] = u
	return u, nil
}

func DeleteUser(id string) error {
	delete(listUser, id)
	return nil
}

func UpdateUser(id string, u2 *model.User) (*model.User, error) {
	u, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	u.FirstName = u2.FirstName
	u.LastName = u2.LastName
	u.Email = u2.Email
	u.Password = u2.Password

	return u, nil
}
