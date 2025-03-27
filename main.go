package main

import (
	"fmt"
	"go-playground/dataStructs"
)

func main() {
	list := dataStructs.NewLinkedList()
	fmt.Println("Создан новый список:", list)

	list.AddLast("один")
	list.AddLast("два")
	list.AddLast("три")
	fmt.Println("После добавления элементов:", list)

	list.AddFirst("ноль")
	fmt.Println("После добавления в начало:", list)

	fmt.Println("Содержит 'два'?", list.Contains("два"))
	fmt.Println("Содержит 'четыре'?", list.Contains("четыре"))

	if val, ok := list.GetAt(2); ok {
		fmt.Println("Элемент по индексу 2:", val)
	}

	if val, ok := list.RemoveFirst(); ok {
		fmt.Println("Удален первый элемент:", val)
		fmt.Println("Список после удаления:", list)
	}

	if val, ok := list.RemoveAt(1); ok {
		fmt.Println("Удален элемент по индексу 1:", val)
		fmt.Println("Список после удаления:", list)
	}

	fmt.Println("Размер списка:", list.GetSize())
}
