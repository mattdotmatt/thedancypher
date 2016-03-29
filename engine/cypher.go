package engine

func Encrypt(message string) string {

	afterStep1 := Step1(message)
	afterStep2 := Step2(afterStep1)
	return Step3(afterStep2)
}
