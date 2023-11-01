package main

type customFunc func(s string)

func (cf customFunc) receiverFunc() {}

func cfFactory() customFunc {
	return func(s string) {}
}

func main() {
	cfFactory().receiverFunc()
}
