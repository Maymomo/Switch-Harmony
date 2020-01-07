package db

import (
	"fmt"
	"testing"
)

func TestQueryDetailByID(t *testing.T) {
	ret,err := QueryDetailByID(2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
