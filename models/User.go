package models

type SwipeEntry struct {
	timestamp string
	time      float32
}

type User struct {
	Username     string
	Creationdate string
	Email        string
	Totaltime    float32
	Swipedata    []SwipeEntry
	ProfilPic    string
}

func GetNullUser() User {
	return User{
		Username:     "",
		Creationdate: "",
		Email:        "",
		Totaltime:    -1.0,
		Swipedata:    nil,
		ProfilPic:    "",
	}
}

func CreateUserObject(username string, creationdate string, email string, totaltime float32, pic string) User {
	newUser := User{
		Username:     username,
		Creationdate: creationdate,
		Email:        email,
		Totaltime:    totaltime,
		Swipedata:    nil,
		ProfilPic:    pic,
	}

	return newUser
}

//func (u *User) SetPassword(password string) {
//}
//
//func (u *User) CheckPassword(password string) bool {
//}
