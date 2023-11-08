package todo

import "testing"

func TestList_Add(t *testing.T) {
	l := List{}
	l.Add("task1")

	if len(l) != 1 {
		t.Errorf("Expected list length of 1, got %d instead.\n", len(l))
	}
}

func TestList_Complete(t *testing.T) {
	l := List{}
	l.Add("task1")
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Expected task to be done, got %v instead.\n", l[0].Done)
	}
}

func TestList_Delete(t *testing.T) {
	l := List{}
	l.Add("task1")
	l.Delete(1)

	if len(l) != 0 {
		t.Errorf("Expected list length of 0, got %d instead.\n", len(l))
	}
}

func TestList_Load(t *testing.T) {
	l := List{}
	l.Add("task1")
	l.Save("test.json")

	l = List{}
	l.Load("test.json")

	if len(l) != 1 {
		t.Errorf("Expected list length of 1, got %d instead.\n", len(l))
	}
}

func TestList_Save(t *testing.T) {
	l := List{}
	l.Add("task1")
	l.Save("test.json")

	l = List{}
	l.Load("test.json")

	if len(l) != 1 {
		t.Errorf("Expected list length of 1, got %d instead.\n", len(l))
	}
}
