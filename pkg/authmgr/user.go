package authmgr

// todo: implement admin user and create user logic

var allowedUsers = map[string]string{
	// username and plain password
	"user.admin": "admin.password",
}

func isAllowedUser(username, password string) bool {
	return allowedUsers[username] == password
}

type AuthOpt struct {
	Username, Password string
}
