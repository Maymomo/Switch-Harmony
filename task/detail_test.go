package task

import (
	"fmt"
	"testing"
)

func TestGetGameDetailInfo(t *testing.T) {
	detail, err := GetGameDetailInfo("http://www.eshop-switch.com/game/2.html")
	fmt.Println(detail, err)
}
