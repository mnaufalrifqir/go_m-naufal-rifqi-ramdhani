package main

//Untuk nama awal Struct gunakan huruf besar, untuk menandakan bahwa itu nama struct bukan variable
type User struct { 
	id       int
	username int
	password int
}
//Untuk nama awal Struct gunakan huruf besar, untuk menandakan bahwa itu nama struct bukan variable
type UserService struct {
	//Gunakan nama variable yang jelas, agar mudah dimengerti
	listUser []User
}
//Untuk setiap pergantian kata pada nama fungsi diawali dengan huruf kapital
//Gunakan nama variable "user" agar mudah dimengerti jangan gunakan "u" saja
func (user UserService) getAllUsers() []User {
	return user.listUser
}
//Untuk setiap pergantian kata pada nama fungsi diawali dengan huruf kapital
//Gunakan nama variable "user" agar mudah dimengerti jangan gunakan "u" saja
func (user UserService) getUserById(id int) User {
	//Gunakan nama variable yang jelas, agar mudah dimengerti
	for _, value := range user.listUser {
		if id == value.id {
			return value
		}
	}

	return User{}
}

// Berapa banyak kekurangan dalam penulisan kode tersebut? 8
// Bagian mana saja terjadi kekurangan tersebut? 
// Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.