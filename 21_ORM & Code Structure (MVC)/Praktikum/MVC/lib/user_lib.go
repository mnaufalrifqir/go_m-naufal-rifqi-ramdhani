package lib

import (
	"mvc_alterra/database"
	"mvc_alterra/middlewares"
	"mvc_alterra/model"
	"mvc_alterra/util"
)

func LoginUsers(user *model.User) (interface{}, error) {
	var err error
	var userDB model.User
	if err = database.DB.First(&userDB, "email = ?", user.Email).Error; err != nil {
		return nil, err
	}

	if err = util.ComparePassword(userDB.Password, user.Password); err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err = database.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
