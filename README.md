## 📚 Library Management System (Go CLI)

A simple command-line Library Management System written in Go (Golang).
This program allows users to manage a collection of books — including adding, viewing, borrowing, returning, searching, saving, and loading book data.
---
## 🚀 Features

- Add Book — Add a new book with title, author, ISBN, and number of copies.

- Show Books — Display all books currently stored in the library.

- Borrow Book — Borrow a book by its ISBN (decreases available copies).

- Return Book — Return a book by its ISBN (increases available copies).

- Search Book — Search by book title or author name.

- Save Data — Save all library data to a JSON file (library.json).

- Load Data — Load existing data from the JSON file.

- Exit — Exit the program safely.

---

```bash
go run main.go

---

# Example Session

Welcome to the Library Management System

> add book
Enter book title: The Go Programming Language
Enter author: Alan Donovan
Enter ISBN: 1001
Enter number of Copies: 3
Books added successfully!

> show books
1. The Go Programming Language - Alan Donovan (1001) | 3

> borrow book
Enter ISBN: 1001
✅ You borrowed: The Go Programming Language
Remaining copies: 2

> save
File Saved.

> exit
Bye Byee 👋


