package main

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
	Copies int
}

type Reader struct {
	Name string
	ID   string
}

type Library struct {
	Books   []Book
	Readers []Reader
	Loans   map[string][]string
}

func (l *Library) AddBook(book Book) {
	for i, b := range l.Books {
		if b.ISBN == book.ISBN {
			l.Books[i].Copies += book.Copies
			fmt.Println("Книга уже существует. Увеличено количество экземпляров.")
			return
		}
	}
	l.Books = append(l.Books, book)
	fmt.Println("Книга добавлена.")
}

func (l *Library) RemoveBook(isbn string) {
	for i, b := range l.Books {
		if b.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Println("Книга удалена.")
			return
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) RegisterReader(reader Reader) {
	for _, r := range l.Readers {
		if r.ID == reader.ID {
			fmt.Println("Читатель уже зарегистрирован.")
			return
		}
	}
	l.Readers = append(l.Readers, reader)
	fmt.Println("Читатель зарегистрирован.")
}

func (l *Library) IssueBook(readerID, isbn string) {
	for i, b := range l.Books {
		if b.ISBN == isbn {
			if b.Copies > 0 {
				l.Books[i].Copies--

				if l.Loans == nil {
					l.Loans = make(map[string][]string)
				}
				l.Loans[readerID] = append(l.Loans[readerID], isbn)
				fmt.Println("Книга выдана читателю.")
				return
			} else {
				fmt.Println("Копий не осталось.")
				return
			}
		}
	}
	fmt.Println("Книга не найдена.")
}

func (l *Library) ReturnBook(readerID, isbn string) {
	loanedBooks, ok := l.Loans[readerID]
	if !ok {
		fmt.Println("У читателя нет книг.")
		return
	}

	for i, loanedISBN := range loanedBooks {
		if loanedISBN == isbn {
			l.Loans[readerID] = append(loanedBooks[:i], loanedBooks[i+1:]...)
			for j, b := range l.Books {
				if b.ISBN == isbn {
					l.Books[j].Copies++
					break
				}
			}
			fmt.Println("Книга возвращена.")
			return
		}
	}
	fmt.Println("Читатель не брал эту книгу.")
}

func main() {
	library := Library{}

	library.AddBook(Book{Title: "1984", Author: "George Orwell", ISBN: "123", Copies: 3})
	library.AddBook(Book{Title: "Clean Code", Author: "Robert C. Martin", ISBN: "456", Copies: 2})

	reader := Reader{Name: "Alihan", ID: "R001"}
	library.RegisterReader(reader)

	library.IssueBook("R001", "123")
	library.IssueBook("R001", "456")

	library.ReturnBook("R001", "123")

	library.RemoveBook("456")
}
