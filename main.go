package main

import (
	"fmt"
	CriminalService "rs-police-db/services"
)

func main() {
	fmt.Println("База Данных Полиции Racoon City")
	const helloMessage = "1. Распечатать Базу Данных\n2. Добавить преступника\n3. Найти преступника по ФИО\n5. Выйти"
	exit := false
	for !exit {
		fmt.Println(helloMessage)
		var w1 string
		fmt.Println("Введите значение: ")
		fmt.Scan(&w1)
		if w1 == "1" {
			fmt.Println("Таблица преступников")
			CriminalService.ShowAllCriminals()
		} else if w1 == "2" {
			var name string
			fmt.Print("Имя: ")
			fmt.Scan(&name)
			fmt.Println()

			var age uint
			fmt.Print("Возраст: ")
			fmt.Scan(&age)
			fmt.Println()

			var gender string
			fmt.Print("Пол: ")
			fmt.Scan(&gender)
			fmt.Println()

			var crimeDate string
			fmt.Print("Дата преступления: ")
			fmt.Scan(&crimeDate)
			fmt.Println()

			var crimeType string
			fmt.Print("Тип преступления: ")
			fmt.Scan(&crimeType)
			fmt.Println()

			var crimeDesc string
			fmt.Print("Описание: ")
			fmt.Scan(&crimeDesc)
			fmt.Println()

			var status string
			fmt.Print("Статус: ")
			fmt.Scan(&status)
			fmt.Println()

			CriminalService.AddCriminalToTable(name, age, gender, crimeDate, crimeType, crimeDesc, status)
		} else if w1 == "3" {
			fmt.Println("Найти преступника по имени")

			var name string
			fmt.Print("Имя: ")
			fmt.Scan(&name)
			fmt.Println()

			fmt.Println(CriminalService.FindByName(name))
		} else if w1 == "4" {
			fmt.Println("Статистическая информация")
		} else if w1 == "5" {
			exit = true
		}
		fmt.Println()
	}
	fmt.Println("Работа программы завершена")

}
