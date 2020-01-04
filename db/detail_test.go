package db

import (
	"fmt"
	"testing"
)

func TestQueryDetailByID(t *testing.T) {
	ret := QueryDetailByID(2)
	fmt.Println(ret)
}
