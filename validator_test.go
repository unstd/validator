package validator

import (
	"fmt"
	"testing"
)

func TestStringHolder_MustHasSuffix(t *testing.T) {
	err := Of("abc").MustString("value 必须是string").Check()
	if err != nil {
		t.Fail()
	}
	err = Of(2).MustString("value 必须是string").Check()
	if err == nil {
		t.Fail()
	}
	err = StringOf("abc").
		MustHasPrefix("xx", "前缀必须是 xx").
		MustHasSuffix("oo", "后缀必须是 oo").
		ThenString("def").
		CheckLength(0, 4, "name长度必须小于3").
		Check()
	fmt.Printf("error message: %s\n", err.Error())
}
