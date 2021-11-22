package daily_practice

type PwdLogin struct {
	s string
}

type EmailLogin struct {
	d int
}

type ILogin interface {
	Login(id uint)
}

func (PwdLogin) Login(id uint) {

}

func (EmailLogin) Login(id uint) {

}

func loginHelper(iLogin ILogin) func(id uint) {
	return iLogin.Login
}

func test() {
	loginHelper(new(EmailLogin))(1)
}

