package users 

type WrongUsernameOrPasswordError struct {}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username of password"
}