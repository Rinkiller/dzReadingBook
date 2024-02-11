// На основе шаблона, напишите url менеджер - программу, сохраняющую ссылки с небольшой
// информацией о них.
// Для решения задачи используйте структуры. Обязательными полями структуры должны быть:
// ● дата добавления;
// ● имя ссылки;
// ● теги;
// ● url.

// ЗАПИСНАЯ КНИЖКА

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var listOfUrl = make([]Item, 0, 0)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func addItem(Name string, Tags string, Link string) Item {
	return Item{Name: Name, Date: time.Now(), Tags: Tags, Link: Link}
}

func printMenu() {
	fmt.Println("")
	fmt.Println("1 - Добавить ссылку")
	fmt.Println("2 - Удалить ссылку")
	fmt.Println("3 - Вывести весь список сохраненных ссылок")
	fmt.Println("7 - Выход из программы")
	fmt.Println("0 - Вывести меню")
	fmt.Println("")
}

func addNewItem() {
	var end bool
	for end != true {
		fmt.Print("Введите имя ссылки:  ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		fmt.Print("Введите теги:  ")
		teg, _ := reader.ReadString('\n')
		fmt.Print("Введите url адрес:  ")
		url, _ := reader.ReadString('\n')
		fmt.Println("name: ", name, " teg: ", teg, " url: ", url)
		for {
			var sel string
			fmt.Print("введенные данные верны y - да / n - нет:  ")
			fmt.Scanf("%s\n", &sel)
			if sel == "y" {
				end = true
				// запись в listOfUrl
				listOfUrl = append(listOfUrl, addItem(strings.TrimRight(name, "\r\n"), strings.TrimRight(teg, "\r\n"), strings.TrimRight(url, "\r\n")))
				break
			} else if sel != "n" {
				fmt.Println("Выберите только y или n")
			}
		}

	}
}

func printListOfUrl() {

	for index := range listOfUrl {
		fmt.Print(index+1, " -  Имя ссылки: ", listOfUrl[index].Name, "\n")
		fmt.Printf("Теги: %s \n", listOfUrl[index].Tags)
		fmt.Printf("Url адрес ссылки: %s \n", listOfUrl[index].Link)
		fmt.Printf("Дата создания записи: %s \n", listOfUrl[index].Date.Format("2006-01-02 15:04:05"))
		fmt.Println()

	}
}

func deleteItemIsList(index int) {
	if index == 0 {
		listOfUrl = listOfUrl[index+1:]
		return
	} else if index == len(listOfUrl)-1 {
		listOfUrl = listOfUrl[0:index]
		return
	}
	listFerst := listOfUrl[0:index]
	listOver := listOfUrl[index+1:]
	listOfUrl = append(listFerst, listOver...)
}

func delitItem() {
	printListOfUrl()
	var index int
	for {
		fmt.Print("Выберите индекс удаляемой записи:  ")
		fmt.Scanf("%d\n", &index)
		if (index > len(listOfUrl)) || (index < 1) {
			fmt.Println("Значение выбранного индекса должно быть в приделах  1 - ", len(listOfUrl))
		} else {
			fmt.Print("Удалить даннуя запись  y - да / n - нет:")
			for {
				var sel string
				fmt.Scanf("%s\n", &sel)
				if sel == "y" {
					deleteItemIsList(index - 1)
					return
				} else if sel != "n" {
					fmt.Println("Выберите только y или n")
				} else {
					return
				}
			}
		}
	}
}

func readBaseFromFile() bool {
	fContent, err := ioutil.ReadFile("base.udb")
	if err != nil {
		return false
	}
	str := string(fContent)[:len(string(fContent))-1]
	words := strings.Split(str, "\n")

	for itm := range words {
		var start, end, step int
		var name, date, tags string
		for index := range words[itm] {
			if words[itm][index] == '{' {
				start = index
			}
			if words[itm][index] == '}' {
				end = index
				step++
				if step == 1 {
					name = words[itm][start+1 : end]
				}
				if step == 2 {
					date = words[itm][start+1 : end]
				}
				if step == 3 {
					tags = words[itm][start+1 : end]
				}
				if step == 4 {
					time, _ := time.Parse("2006-01-02 15:04:05", date)
					step = 0
					listOfUrl = append(listOfUrl, Item{Name: name, Date: time, Tags: tags, Link: words[itm][start+1 : end]})

				}

			}

		}

	}
	return true
}
func saveBaseInFile() {
	f, err := os.Create("base.udb")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := range listOfUrl {
		_, err = f.WriteString("{" + string(listOfUrl[i].Name) + "}" + "{" + listOfUrl[i].Date.Format("2006-01-02 15:04:05") + "}" + "{" + string(listOfUrl[i].Tags) + "}" + "{" + string(listOfUrl[i].Link) + "}\n")
		if err != nil {
			panic(err)
		}
	}
}

func nullListDB() {
	fmt.Println("В базе нет ни одной записи!!!")
	fmt.Print("Давайте добавим запись  y - да / n - нет:")
	for {
		var sel string
		fmt.Scanf("%s\n", &sel)
		if sel == "y" {
			addNewItem()
			break
		} else if sel != "n" {
			fmt.Println("Выберите только y или n")
		} else {
			fmt.Println("Спасибо за работу")
			os.Exit(1) // окончание работы
		}
	}
}

func main() {
	if readBaseFromFile() != true {
		fmt.Println("DEBUGING!!!!!")
		nullListDB()
		saveBaseInFile()
	}
	for {
		if len(listOfUrl) == 0 { // В базе нет ни одной записи
			nullListDB()
			saveBaseInFile()
		}
		var pos int
		fmt.Println("Выберите пункт меню")
		fmt.Println("0 - Вывести меню")
		fmt.Print("Ваш выбор:    ")
		fmt.Scanf("%d\n", &pos)

		switch pos {

		case 0:
			printMenu()
		case 1:
			{ // Добавляем запись           имя ссылки;  теги;  url.
				addNewItem()
				saveBaseInFile()
			}

		case 2:
			{ // удалить запись
				delitItem()
				saveBaseInFile()
			}
		case 3:
			printListOfUrl()

		case 4:

		case 5:

		case 6:

		case 7:
			{
				fmt.Println("Спасибо за работу")
				return
			}
		default:
			fmt.Println("Выберите какой либо пункт меню")

		}
	}

}
