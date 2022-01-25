package tree

import (
	"fmt"
)

type Element struct {
	Value int
}

type Tree struct {
	Element *Element
	right   *Tree
	left    *Tree
}

func (t *Tree) IsEmpty() bool {

	return t.Element == nil
}

func (tree *Tree) Add(newElement *Element) {

	//Esta vazio inicia a arvore
	if tree.IsEmpty() {

		tree.Element = newElement
	} else {

		newTree := Tree{
			Element: newElement,
		}

		if newElement.Value < tree.Element.Value { //Add a esquerda

			if tree.left == nil { // é um nó folha

				tree.left = &newTree
			} else {
				tree.left.Add(newElement) //Repassa a responsabilidade
			}

		} else if newElement.Value > tree.Element.Value { //add na descendencia direita

			if tree.right == nil {

				tree.right = &newTree
			} else {
				tree.right.Add(newElement)
			}
		}
	}
}

func (t *Tree) Find(value int) bool {

	if t.IsEmpty() {
		return false
	}

	if t.Element.Value == value {
		return true
	} else {
		if value < t.Element.Value {

			if t.left == nil {
				return false
			} else {
				return t.left.Find(value)
			}
		} else if value > t.Element.Value {

			if t.right == nil {
				return false
			} else {
				return t.right.Find(value)
			}
		}

		return false
	}
}

func (t *Tree) Remove(element *Element) *Tree {

	//achou elemento
	if t.Element.Value == element.Value {
		//caso elemento esta em unico nó
		if t.right == nil && t.left == nil {
			return nil
		} else {
			//tem nodes a esquerda e não a direita
			if t.left != nil && t.right == nil {
				return t.left
				//tem nodes a direita e não a esquerda
			} else if t.right != nil && t.left == nil {
				return t.right
				//existem nodes tanto a esquerda como na direita
			} else {
				treeAux := t.left
				for treeAux.right != nil {
					treeAux = treeAux.right
				}
				//troca elementos da tree
				t.Element = treeAux.Element
				treeAux.Element = element
				t.left = t.left.Remove(element)
			}
		}

	} else if element.Value < t.Element.Value {
		t.left = t.left.Remove(element) //chama a esquerda
	} else if element.Value > t.Element.Value {
		t.right = t.right.Remove(element) //chama a direita
	}

	return t
}

func (t *Tree) ViewPreOrdem() {

	if !t.IsEmpty() {

		fmt.Print(t.Element.Value, " ")

		if t.left != nil {
			t.left.ViewPreOrdem()
		}

		if t.right != nil {
			t.right.ViewPreOrdem()
		}
	}
}

func (t *Tree) ViewInOrdem() {

	if !t.IsEmpty() {

		if t.left != nil {
			t.left.ViewInOrdem()
		}

		fmt.Print(t.Element.Value, " ")

		if t.right != nil {
			t.right.ViewInOrdem()
		}
	}
}

func (t *Tree) ViewPosOrdem() {

	if !t.IsEmpty() {

		if t.right != nil {
			t.right.ViewPosOrdem()
		}

		if t.left != nil {
			t.left.ViewInOrdem()
		}

		fmt.Print(t.Element.Value, " ")
	}
}

func (t *Tree) ViewInvertedOrdem() {

	if !t.IsEmpty() {

		if t.right != nil {
			t.right.ViewInvertedOrdem()
		}

		fmt.Print(t.Element.Value, " ")

		if t.left != nil {
			t.left.ViewInvertedOrdem()
		}
	}
}
