package tree

import (
	"testing"
)

func TestTree_IsEmptyWhenNewTree(t *testing.T) {

	tree := Tree{}

	value := tree.IsEmpty()

	if !value {
		t.Error("Value should be True")
	}
}

func TestTree_IsNotEmptyWhenNewTreeInicialized(t *testing.T) {

	tree := Tree{
		Element: &Element{10},
	}

	value := tree.IsEmpty()

	if value {
		t.Error("Value should be NotEmpty")
	}
}

func TestTree_FindNumberExistsInTree(t *testing.T) {

	tree := Tree{
		Element: &Element{10},
	}

	tree.Add(&Element{5})
	tree.Add(&Element{1})
	tree.Add(&Element{8})
	tree.Add(&Element{12})
	tree.Add(&Element{18})

	value := tree.Find(5)

	if !value {
		t.Error("Value 5 should exists")
	}
}

func TestTree_FindNumberNotExistsInTree(t *testing.T) {

	tree := Tree{
		Element: &Element{10},
	}

	tree.Add(&Element{5})
	tree.Add(&Element{1})
	tree.Add(&Element{8})
	tree.Add(&Element{12})
	tree.Add(&Element{18})

	value := tree.Find(52)

	if value {
		t.Error("Value 52 should not exists")
	}
}

func TestTree_Remove(t *testing.T) {

	tree := Tree{
		Element: &Element{10},
	}

	tree.Add(&Element{5})
	tree.Add(&Element{1})
	tree.Add(&Element{8})
	tree.Add(&Element{12})
	tree.Add(&Element{18})

	newTree := tree.Remove(&Element{8})
	newTree = tree.Remove(&Element{18})

	value := newTree.Find(8)
	value2 := newTree.Find(18)

	if value {
		t.Error("Value 8 should be removed")
	}

	if value2 {
		t.Error("Value 18 should be removed")
	}
}

func TestTree_ViewPreOrdem(t *testing.T) {

	expected := "10 5 1 8 15 12 18"

	value := ""

	if expected != value {
		t.Error("ViewInOrdem should be", "10 5 1 8 15 12 18")
	}
}

func TestTree_ViewInOrdem(t *testing.T) {

	expected := "1 5 8 10 12 15 18"

	value := ""

	if expected != value {
		t.Error("ViewInOrdems hould be", "1 5 8 10 12 15 18")
	}
}

func TestTree_ViewPosOrdem(t *testing.T) {

	expected := "18 12 15 1 5 8 10"

	value := ""

	if expected != value {
		t.Error("ViewPosOrdem should be", "18 12 15 1 5 8 10")
	}
}

func TestTree_ViewInvertedOrdem(t *testing.T) {

	expected := "18 15 12 10 8 5 1"

	value := ""

	if expected != value {
		t.Error("ViewInvertedOrdem should be", "18 15 12 10 8 5 1")
	}
}
