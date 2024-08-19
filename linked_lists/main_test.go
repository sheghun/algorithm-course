package main

import "testing"

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[string]()

	list.Append("2")
	list.Append("3")

	t.Run("Check Length", func(t *testing.T) {
		if list.Len() != 2 {
			t.Fatalf("length not increasing")
		}
	})

	t.Run("Get with index", func(t *testing.T) {
		if *list.Get(1) != "3" {
			t.Fatalf("getting with index failed")
		}
	})

	t.Run("InsertAt index", func(t *testing.T) {
		list.InsertAt("4", 1)
		if *list.Get(1) != "4" || list.Len() != 3 || *list.Get(2) != "3" {
			t.Fatalf("inserting at index failed")
		}
	})

	t.Run("Prepend item", func(t *testing.T) {
		list.Prepend("0")
		if *list.Get(0) != "0" || *list.Get(1) != "2" || list.Len() != 4 {
			t.Fatalf("prepending failed")
		}

		list.Prepend("Z")
		if *list.Get(0) != "Z" || *list.Get(1) != "0" || list.Len() != 5 {
			t.Fatalf("prepending failed")
		}
	})

	t.Run("Append item", func(t *testing.T) {
		list.Append("5")
		if *list.Get(list.Len() - 1) != "5" || list.Len() != 6 {
			t.Fatalf("appending item failed")
		}
	})

	t.Run("remove at", func(t *testing.T) {
		val := *list.Get(4)
		val1 := *list.RemoveAt(4)
		if val != val1 || list.Len() != 5 {
			t.Fatalf("expected: %s, got: %s", val, val1)
		}

		if val1 == *list.Get(4) {
			t.Fatalf("expected: %s, got %s", *list.Get(5), val)
		}
	})

	t.Run("remove using item", func(t *testing.T) {
		val := *list.Remove("5")
		if val != "5" || list.Len() != 4 {
			t.Fatalf("expected: %s", *list.Get(list.Len() - 1))
		}

		if val == *list.Get(list.Len() - 1) {
			t.Fatalf("removing item failed")
		}
	})
}
