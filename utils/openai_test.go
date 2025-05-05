package utils

import (
	"fmt"
	"testing"
)

func TestThinkFile(t *testing.T) {
	msg, err := ThinkFile("../uploads/fd019793a01ef28dd8980cc0d8a1fa94")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(msg)
}
