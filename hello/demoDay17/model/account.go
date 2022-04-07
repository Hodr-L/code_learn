package model

type good struct {
	pace string
}

type account struct {
	good
	username string
	password string
	sal      float64
}

func User(name string, pwd string, sal float64) *account {
	return &account{
		username: name,
		password: pwd,
		sal:      sal,
	}
}
func (a *account) Setname(name string) {
	a.username = name
}
func (a *account) Setpwd(pwd string) {
	a.password = pwd
}
func (a *account) Setsal(sal float64) {
	a.sal = sal
}
func (a *account) Getname() string {
	return a.username
}
func (a *account) Getpwd() string {
	return a.password
}
func (a *account) Getsal() float64 {
	return a.sal
}
