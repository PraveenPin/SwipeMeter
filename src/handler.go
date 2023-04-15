package src

type user struct {
	email          string
	username       string
	passwordhash   string
	createDate     string
	totalTime      float32
	profilePicLink string
}

func GetUserObject(username string) (user, bool) {
	userList := [1]user{
		user{
			"praveen.pinjala@rutgers.edu",
			"praveenpin",
			"123",
			"2023/4/14",
			100.0,
			"",
		},
	}
	//get userList from database
	for _, user := range userList {
		if user.username == username {
			return user, true
		}
	}
	return user{}, false
}

func (u *user) validatePasswordHash(passwd string) bool {
	return u.passwordhash == passwd
}

func addUserObject(email string, username string, passwd string) bool {
	newUser := user{
		email,
		username,
		passwd,
		"",
		0,
		"",
	}

	userList := []user{
		user{
			"praveen.pinjala@rutgers.edu",
			"praveenpin",
			"123",
			"2023/4/14",
			100.0,
			"",
		},
	}

	for _, ele := range userList {
		if ele.email == email || ele.username == username {
			return false
		}
	}

	userList = append(userList, newUser)
	return true
}
