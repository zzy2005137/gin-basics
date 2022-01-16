package util

import (
	"fmt"
	"testing"
)

func Test_GenerateJwt(t *testing.T) {

	fmt.Println(GenerateJwt("zzy", "1230"))
	fmt.Println(GenerateJwt("sss", "1230"))
	// fmt.Println("hello")
}

func Test_ParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyTmFtZSI6Inp6eSIsIlBhc3N3b3JkIjoiMTIzMCIsImV4cCI6MTY0MjM0NTk4OCwiaXNzIjoiZ2luLWV4YW1wbGUifQ.eix8x_W4HYXDRpwGhVgCXvwc620tCU2n6Gd0coUnXB4"
	_, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	println("success, c.Next()")
}
