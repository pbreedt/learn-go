package generics

import (
	"testing"
)

func setupLinkedListString() *LinkedList[string] {
	var linkList *LinkedList[string]
	linkList = linkList.Append("Hello")
	linkList = linkList.Append("Hola")
	linkList = linkList.Append("हैलो")

	return linkList
}

func TestLinkedListLen(t *testing.T) {
	linkList := setupLinkedListString()
	want := 3
	got := linkList.Len()
	if want != got {
		t.Errorf("incorrect result. wanted:%d, got:%d", want, got)
	}
}

func TestLinkedListString(t *testing.T) {
	linkList := setupLinkedListString()
	want := "Hello->Hola->हैलो->nil"
	got := linkList.String()
	if want != got {
		t.Errorf("incorrect result. wanted:%s, got:%s", want, got)
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	linkList := setupLinkedListString()
	linkList.InsertAt(2, "Bye")
	want := "Hello->Hola->Bye->हैलो->nil"
	got := linkList.String()
	if want != got {
		t.Errorf("incorrect result. wanted:%s, got:%s", want, got)
	}
}

func TestLinkedListAppend(t *testing.T) {
	linkList := setupLinkedListString()
	linkList.Append("Bye")
	linkList.Append("Adiós")
	want := "Hello->Hola->हैलो->Bye->Adiós->nil"
	got := linkList.String()
	if want != got {
		t.Errorf("incorrect result. wanted:%s, got:%s", want, got)
	}
}

func TestLinkedListContains(t *testing.T) {
	linkList := setupLinkedListString()
	want := true
	got := linkList.Contains("Hello")
	if want != got {
		t.Errorf("incorrect result. wanted:%t, got:%t", want, got)
	}

	want = false
	got = linkList.Contains("Bye")
	if want != got {
		t.Errorf("incorrect result. wanted:%t, got:%t", want, got)
	}
}

type Person struct {
	Name string
	Age  int
}

func TestLinkedListCustomType(t *testing.T) {
	var peopleList *LinkedList[Person]
	peopleList = peopleList.Append(Person{"Fred", 23})
	peopleList = peopleList.Append(Person{"Joan", 30})

	want := "{Fred 23}->{Joan 30}->nil"
	got := peopleList.String()
	if want != got {
		t.Errorf("incorrect result. wanted:%s, got:%s", want, got)
	}

	want2 := true
	got2 := peopleList.Contains(Person{"Joan", 30})
	if want2 != got2 {
		t.Errorf("incorrect result. wanted:%t, got:%t", want2, got2)
	}

	want3 := false
	got3 := peopleList.Contains(Person{"Joan", 31})
	if want3 != got3 {
		t.Errorf("incorrect result. wanted:%t, got:%t", want3, got3)
	}
}
