package structs

import (
	"reflect"
	"testing"
)

func TestNewCappedStringCollection(t *testing.T) {
	a:=NewCappedStringCollection(4)

	t.Log(a.AddString("aaaa1"))
	t.Log(a.AddString("aaaa2"))
	t.Log(a.AddString("aaaa3"))
	t.Log(a.AddString("aaaa4"))
	t.Log(a.AddString("aaaa5"))
	t.Log(a.AddString("aaaa6"))
	t.Log(a.AddString("aaaa7"))
	t.Log(a.AddString("aaaa8"))
	t.Log(a.AddString("aaaa9"))
	t.Log(a.AddString("aaaa10"))
	if !reflect.DeepEqual(a.GetSavedStrings(), []string{"aaaa7", "aaaa8", "aaaa9", "aaaa10"}){
		t.Fail()
	}
	t.Log(a.GetSavedStrings())
}