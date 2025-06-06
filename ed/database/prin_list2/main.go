package main

import (
	"container/list"
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	elem := "["
	for i := l.Front(); i != nil; i = i.Next() {
		if i == sword {
			if i.Value.(int) > 0 {
				elem += fmt.Sprintf(" %d>",i.Value.(int))
			} else {
				elem += fmt.Sprintf(" <%d", i.Value.(int))
			}
		} else {
			elem +=fmt.Sprintf(" %d",i.Value.(int))
		}
	}
	elem += " ]"
	return elem
}

// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
	if it.Next() == nil{
		return l.Front()
	}
	return it.Next()
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
	if it.Prev() == nil{
		return l.Back()
	}
	return it.Prev()
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}
	for l.Len() > 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			toRemove := Next(l, sword)
			l.Remove(toRemove)
			sword = Next(l, sword)
		} else {
			toRemove := Prev(l, sword)
			l.Remove(toRemove)
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}
