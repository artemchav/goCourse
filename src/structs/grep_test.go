package structs

import (
	"fmt"
	"reflect"
	"testing"
)

// Еще один вид тестов. Вывод в консоль должен совпадать с тем, что написано
// в комментариях после Output. Тест пройден, если вывелось то же самое.
// Кроме того, этот пример будет отображен на странице GoDoc.
func ExampleNewMatchesStorage() {
	afterCount := 4
	matches := NewMatchesStorage(afterCount)

	strings := []string{
		"asdfasf1",
		"asdfasf2",
		"asdfasf3",
		"asdfasf4",
		"asdfasf5",
		"asdfasf6",
		"asdfasf7",
		"asdfasf8",
		"asdfasf9",
		"asdfasf10",
		"asdfasf11",
		"asdfasf12",
	}
	capped := NewCappedStringCollection(3)
	for i, v := range strings {
		before:=capped.GetSavedStrings()
		capped.AddString(v)
		res := matches.Process(v, i%3 == 0, before)
		for _, m:=range res{
			fmt.Println(m.Result())
		}
	}
	res := matches.Finish()
	for _, m:=range res{
		fmt.Println(m.Result())
	}
	// Output:
	// [asdfasf1 asdfasf2 asdfasf3 asdfasf4 asdfasf5]
	// [asdfasf1 asdfasf2 asdfasf3 asdfasf4 asdfasf5 asdfasf6 asdfasf7 asdfasf8]
	// [asdfasf4 asdfasf5 asdfasf6 asdfasf7 asdfasf8 asdfasf9 asdfasf10 asdfasf11]
	// [asdfasf7 asdfasf8 asdfasf9 asdfasf10 asdfasf11 asdfasf12]
}

func TestNewMatchesStorage(t *testing.T) {

	afterCount := 4
	matches := NewMatchesStorage(afterCount)

	strings := []string{
		"asdfasf1",
		"asdfasf2",
		"asdfasf3",
		"asdfasf4",
		"asdfasf5",
		"asdfasf6",
		"asdfasf7",
		"asdfasf8",
		"asdfasf9",
		"asdfasf10",
		"asdfasf11",
		"asdfasf12",
	}
	capped := NewCappedStringCollection(3)
	for i, v := range strings {
		before:=capped.GetSavedStrings()
		capped.AddString(v)
		res := matches.Process(v, i%3 == 0, before)
		for _, m:=range res{
			t.Log("Matches:", m.Result())
		}
	}
	res := matches.Finish()
	for _, m:=range res{
		t.Log("Matches:", m.Result())
	}
}

func TestMatch_AddAfter(t *testing.T) {
	match := &Match{
		IsDone:     false,
		afterCount: 4,
		matched:    "matchedString",
		before: []string{
			"asdfasf1",
			"asdfasf2",
		},
		after: make([]string, 0, 4),
	}
	match.AddAfter("123")
	match.AddAfter("456")
	match.AddAfter("4567889")
	match.AddAfter("45678891098")
	match.AddAfter("09877654")
	match.AddAfter("999999999")
	if !reflect.DeepEqual(match.Result(), []string{"asdfasf1", "asdfasf2", "matchedString", "123", "456", "4567889", "45678891098"}) {
		t.Fail()
	}
	t.Log(match.Result())
}
