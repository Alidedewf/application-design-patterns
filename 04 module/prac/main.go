package main

import "fmt"

type Document interface {
	Open()
}

type Report struct{}

func (r Report) Open() {
	fmt.Println("Открыт документ: Отчет")
}

type Resume struct{}

func (r Resume) Open() {
	fmt.Println("Открыт документ: Резюме")
}

type Letter struct{}

func (l Letter) Open() {
	fmt.Println("Открыт документ: Письмо")
}

type Invoice struct{}

func (i Invoice) Open() {
	fmt.Println("Открыт документ: Счет-фактура")
}

type DocumentCreator interface {
	CreateDocument() Document
}

type ReportCreator struct{}

func (rc ReportCreator) CreateDocument() Document {
	return Report{}
}

type ResumeCreator struct{}

func (rc ResumeCreator) CreateDocument() Document {
	return Resume{}
}

type LetterCreator struct{}

func (lc LetterCreator) CreateDocument() Document {
	return Letter{}
}

type InvoiceCreator struct{}

func (ic InvoiceCreator) CreateDocument() Document {
	return Invoice{}
}

func GetCreatorByType(docType string) DocumentCreator {
	switch docType {
	case "report":
		return ReportCreator{}
	case "resume":
		return ResumeCreator{}
	case "letter":
		return LetterCreator{}
	case "invoice":
		return InvoiceCreator{}
	default:
		return nil
	}
}

func main() {
	var input string

	fmt.Println("Введите тип документа (report, resume, letter, invoice):")
	fmt.Scanln(&input)

	creator := GetCreatorByType(input)
	if creator == nil {
		fmt.Println("Неизвестный тип документа!")
		return
	}

	document := creator.CreateDocument()
	document.Open()
}
