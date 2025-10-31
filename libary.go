package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"encoding/json"
)

type Books struct {
	Name string
	Author string
	ISBN int
	Copies int
}

func main() {
	reader:= bufio.NewReader(os.Stdin)
	var books [] Books

	fmt.Println("\nWelcome to the Libary Managment System")
	fmt.Println(`
- add book
- show books
- borrow book
- return book
- search book
- save 
- load 
- exit
	`)

	for {
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "add book" {
			fmt.Print("Enter book title: ")
			
			name,_ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
//Checking not to add same book name twice
			exists := false

			for _, item := range books {
				if strings.EqualFold(name,item.Name) {
					exists = true
				}
			}
			if exists {
				fmt.Println("Book already existed. Enter another one.")
				continue
			}

			fmt.Print("Enter author: ")

			author,_ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Enter ISBN: ")

			isbnStr,_ := reader.ReadString('\n')
			isbnStr = strings.TrimSpace(isbnStr)

			isbn, err := strconv.Atoi(isbnStr)
			if err != nil {
				fmt.Println("Error. Enter ISBN number.")
				
			}

			fmt.Print("Enter number of Copies: ")

			copiesStr,_ := reader.ReadString('\n')
			copiesStr = strings.TrimSpace(copiesStr)

			copies, err := strconv.Atoi(copiesStr)
			if err != nil {
				fmt.Println("Error. Enter copies number.")
				
			}

			books = append(books, Books{name, author, isbn, copies})
			fmt.Println("Books added succesfully!")

		} else if input == "show books" {
			if len(books) == 0 {
				fmt.Println("Libary is empty.")

			} else{
				fmt.Println("\nLibary Collection:")
				for i,e := range books{
					fmt.Printf("%d. %s - %s (%d) | %d\n",i+1, e.Name, e.Author, e.ISBN, e.Copies)
				}
			}
		
		} else if input == "search book" {
			if len(books) == 0 {
				fmt.Println("Books libary empty.")
			} else {
				fmt.Print("Enter book name or author name: ")

				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)

				found := false

				for _, item := range books {
					if strings.EqualFold(name,item.Name) || strings.EqualFold(name,item.Author) {
						found = true
						fmt.Printf("%s - %s (%d) | %d",item.Name,item.Author, item.ISBN,item.Copies)
						return
					}
				if !found {
					fmt.Printf("Book %s not found",name)
					return
				}
				}
			}
		} else if input == "borrow book" {
			fmt.Print("Enter ISBN: ")

			isbnStr ,_ := reader.ReadString('\n')
			isbnStr = strings.TrimSpace(isbnStr)

			isbn, err := strconv.Atoi(isbnStr)
			if err != nil {
			fmt.Println("Error. Enter correct ISBN number.")
			}
			
			found := false 

			for i := range books {
				if isbn == books[i].ISBN {
					found = true

				if books[i].Copies > 0 {
					books[i].Copies --
					fmt.Printf("✅ You borrowed: %s\n",books[i].Name)
					fmt.Printf("Reamining books: %d\n",books[i].Copies)

				} else {
					fmt.Println("❌ No copies left for %s\n", books[i].Name)
				}

				
				
				}
			}
			if !found {
				fmt.Println("Error. ISBN number not found.")
			}
		} else if input == "return book" {
			fmt.Print("Enter ISBN: ")

			isbnStr,_ := reader.ReadString('\n')
			isbnStr = strings.TrimSpace(isbnStr)

			isbn,err := strconv.Atoi(isbnStr)
			if err != nil {
				fmt.Println("Error. Please enter correct ISBN number.")
			}
			
			found := false

			for i := range books {
				if isbn == books[i].ISBN {
					found = true

				if books[i].Copies >= 0 {
					books[i].Copies ++ 
					fmt.Printf("✅ You returned: %s", books[i].Name)
					fmt.Printf("Available copies: %d", books[i].Copies)
				}

				
				}
			}
			if !found{
				fmt.Println("ISBN number not found.")
			}

		} else if input == "save" {
			file, err := os.Create("libary.json")
	if err != nil {
		fmt.Println("Error saving file.")
	}

	defer file.Close()

	data, err := json.MarshalIndent(books, "", " ") 
	if err != nil {
		fmt.Println("Error converting txt. to json.")
	}
	file.Write(data)
	fmt.Println("Filed Saved.")
		} else if input == "load" {
			data, err := os.ReadFile("libary.json")
	if err != nil {
		fmt.Println("Error loading file.",err)
	}
	
	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Println("Error parsing JSON",err)
	}
	fmt.Println("File loaded succesfully.")
		} else if input == "exit" {
			fmt.Println("\nBye Byee\n")
			break
		} else {
			fmt.Println("Error. Please enter correct operation.")
		}


		fmt.Println(`
	- add book
	- show books
	- borrow book
	- return book
	- search book
	- save 
	- load 
	- exit
		`)
		
	} 
}


