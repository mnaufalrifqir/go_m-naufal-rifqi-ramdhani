package User

func CreateUserController(user User) User {
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return user
}

func DeleteUserController(id int) {
	users = append(users[:id-1], users[id:]...)
}

func UpdateUserController(user User, id int) {
	users[id-1].Name = user.Name
	users[id-1].Email = user.Email
	users[id-1].Password = user.Password
}

func GetUserController(id int) (user User) {
	if id != 0 {
		user = users[id-1]
	}
	return user
}
