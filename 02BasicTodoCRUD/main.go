package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"simple_crud/db"
	"strings"
)

func printMsg(message string) {
	fmt.Println(message)
}

func scanText() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(line)
}

func main() {
	var option int
	todoDB := db.DBConn()
	defer todoDB.Close()

	todoDB.MaybeCreateTable()

	printMsg("# Welcome to Todo app")
	println()

	for {

		printMsg("What you want to do ?")
		printMsg(" 1. Add Todo")
		printMsg(" 2. Get Todo by id")
		printMsg(" 3. Retrive All Todos")
		printMsg(" 4. Update Todo")
		printMsg(" 5. Remove Todo")
		printMsg(" 6. Exit")

		fmt.Print("\nChoose: ")
		fmt.Scanln(&option)

		if option == 6 {
			printMsg("\nBye bye 👋")
			break
		}

		switch option {
		case 1:
			var title, description string

			fmt.Print("Title Please: ")
			title = scanText()

			fmt.Print("Description Please: ")
			description = scanText()

			todoDB.AddTodo(title, description)
		case 2:
			var id int

			fmt.Print("Enter Id: ")
			fmt.Scanln(&id)

			todoDB.GetById(id)
		case 3:
			printMsg("All Todos:")
			todoDB.GetAll()
		case 4:
			var id int
			var title, description string

			fmt.Print("Todo Id Please: ")
			fmt.Scanln(&id)

			todoId := todoDB.GetById(id)

			if todoId > 0 {
				fmt.Print("Title Please: ")
				title = scanText()

				fmt.Print("Description Please: ")
				description = scanText()

				todoDB.UpdateTodo(id, title, description)
			} else {
				printMsg("Sorry Id not found")
			}
		case 5:
			var id int
			fmt.Print("Todo id to remove: ")
			fmt.Scanln(&id)

			todoDB.Remove(id)
		default:
			printMsg("Please select valid option")
		}

	}

}
