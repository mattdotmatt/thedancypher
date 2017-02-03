package engine

func Encrypt(message string) string {
	return Step3(Step2(Step1(message)))
}
