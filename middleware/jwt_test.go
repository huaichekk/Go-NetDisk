package middleware

import (
	"fmt"
	"go-netdisk/models"
	"testing"
)

func TestGetToken(t *testing.T) {
	u := models.User{
		Username: "admin",
		Password: "320930",
		IsVIP:    true,
	}
	token, err := GetToken(u)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)

	parseToken, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(parseToken)
}
