package utils

import "math/rand"

type UserOtp struct {
	Otp string `json:"otp"`
}

func RandNumeric(n int) string {

	// TODO: Generate Random string

	var numeric = []rune("01234567899")

	s := make([]rune, n)
	for i := range s {
		s[i] = numeric[rand.Intn(len(numeric))]
	}
	return string(s)

}
