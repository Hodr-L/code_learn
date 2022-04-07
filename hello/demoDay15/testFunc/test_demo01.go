package testFunc

type MethodUtils struct {
	Name string
}

func (a *MethodUtils) Test(m, n int) int {
	return m * n
}
