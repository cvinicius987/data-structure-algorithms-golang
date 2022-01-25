package tree_avl

import (
	"fmt"
	"strings"
)

type Element struct {
	Value int
}

type Tree struct {
	Element *Element
	right   *Tree
	left    *Tree
	balance int
}

func (t *Tree) IsEmpty() bool {

	return t.Element == nil
}

func (tree *Tree) Add(newElement *Element) *Tree {

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
				tree.left = tree.left.Add(newElement) //Repassa a responsabilidade
			}

		} else if newElement.Value > tree.Element.Value { //add na descendencia direita

			if tree.right == nil {

				tree.right = &newTree
			} else {
				tree.right = tree.right.Add(newElement)
			}
		}
	}

	return tree
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

func (t *Tree) PrintTree(level int) string {

	builder := strings.Builder{}

	builder.WriteString(t.ToString())
	builder.WriteString("\n")

	for i := 0; i < level; i++ {
		builder.WriteString("\t")
	}
	if t.left != nil {
		builder.WriteString("+-ESQ: ")
		builder.WriteString(t.left.PrintTree(level + 1))
	} else {
		builder.WriteString("+-ESQ: NULL")
	}
	builder.WriteString("\n")

	for i := 0; i < level; i++ {
		builder.WriteString("\t")
	}
	if t.right != nil {
		builder.WriteString("+-DIR: ")
		builder.WriteString(t.right.PrintTree(level + 1))
	} else {
		builder.WriteString("+-DIR: NULL")
	}
	builder.WriteString("\n")

	return builder.String()
}

func (t *Tree) ToString() string {
	return fmt.Sprintf("[%d] (%d)", t.Element.Value, t.balance)
}

func (t *Tree) Balance() int {
	return t.balance
}

func (t *Tree) CalculeHeight() int {

	if t.left == nil && t.right == nil {
		return 1
	} else if t.left != nil && t.right == nil {
		return 1 + t.left.CalculeHeight()
	} else if t.left == nil && t.right != nil {
		return 1 + t.right.CalculeHeight()
	} else {

		var result int

		if t.left.CalculeHeight() > t.right.CalculeHeight() {
			result = t.left.CalculeHeight()
		} else {
			result = t.right.CalculeHeight()
		}

		return 1 + result
	}
}

func (t *Tree) CalculeBalance() {

	if t.left == nil && t.right == nil {
		t.balance = 0
	} else if t.left == nil && t.right != nil {
		t.balance = t.right.CalculeHeight() - 0
	} else if t.left != nil && t.right == nil {
		t.balance = t.left.CalculeHeight() - 0
	} else {
		t.balance = t.right.CalculeHeight() - t.left.CalculeHeight()
	}

	if t.right != nil {
		t.right.CalculeBalance()
	}

	if t.left != nil {
		t.left.CalculeBalance()
	}
}

func (t *Tree) VerifyBalance() *Tree {

	if t.balance >= 2 || t.balance <= -2 {
		if t.balance >= 2 {
			if t.balance*t.right.balance > 0 {
				return t.SimpleRotationRight()
			} else {
				return t.DoubleRotationRight()
			}
		} else {
			if t.balance*t.left.balance > 0 {
				return t.SimpleRotationLeft()
			} else {
				return t.DoubleRotationLeft()
			}
		}
	}
	t.CalculeBalance()
	if t.left != nil {
		t.left = t.left.VerifyBalance()
	}
	if t.right != nil {
		t.right = t.right.VerifyBalance()
	}
	return t
}

func (t *Tree) SimpleRotationRight() *Tree {
	var childRight *Tree
	var childChildRight *Tree = nil

	childRight = t.right
	if t.right != nil {
		if t.right.left != nil {
			childChildRight = childRight.left
		}
	}
	childRight.left = t
	t.right = childChildRight
	return childRight
}

func (t *Tree) DoubleRotationRight() *Tree {
	var tree = t
	var childRight = t.right
	var childChild = childRight.left
	var newInserted = childChild.right

	//1º
	childRight.left = newInserted
	childChild.right = childRight
	t.right = childChild

	//2º
	var newChildRight = t.right
	tree.right = nil
	newChildRight.left = tree

	return newChildRight
}

func (t *Tree) SimpleRotationLeft() *Tree {
	var childLeft *Tree
	var childChildLeft *Tree = nil

	childLeft = t.left
	if t.left != nil {
		if t.left.right != nil {
			childChildLeft = childLeft.right
		}
	}
	childLeft.right = t
	t.left = childChildLeft
	return childLeft

}

func (t *Tree) DoubleRotationLeft() *Tree {
	var tree = t
	var childLeft = t.left
	var childChild = childLeft.right
	var newInserted = childChild.left

	//1º
	childLeft.right = newInserted
	childChild.left = childLeft
	t.left = childChild

	//2º
	var newChildLeft = t.left
	tree.left = nil
	newChildLeft.right = tree

	return newChildLeft
}
