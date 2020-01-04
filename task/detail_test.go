package task

import (
	"fmt"
	"testing"
)

func TestGetGameDetailInfo(t *testing.T) {
	detail, err := GetGameDetailInfo("http://www.eshop-switch.com/game/3005.html")
	detail.InsertToDB()
	fmt.Println(detail, err)
}
