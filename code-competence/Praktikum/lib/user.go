package lib

import (
	"code_competence/database"
	"code_competence/middlewares"
	"code_competence/model"
	"code_competence/util"
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

	Token, err := middlewares.CreateToken(userDB.ID)
	if err != nil {
		return nil, err
	}

	return Token, nil
}
