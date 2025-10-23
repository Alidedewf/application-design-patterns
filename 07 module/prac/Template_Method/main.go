package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Logger struct{}

func (l *Logger) Log(step string) {
	log.Printf("[журнал] шаг: %s\n", step)
}

type IReportGenerator interface {
	generateHeader() (string, error)
	formatData(commonData string) (string, error)
	generateFooter() (string, error)

	shouldAddOptionalData() bool
	addOptionalData() (string, error)

	shouldSave() bool
	save(data string) error

	shouldSendEmail() bool
	sendEmail(data string) error
}

func loadBaseData(logger *Logger) (string, error) {
	logger.Log("загрузка общих данных из базы данных...")
	return "общие данные компании", nil
}

func askUser(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt + " (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "y" {
			return true
		}
		if answer == "n" {
			return false
		}
		fmt.Println("некорректный ввод. пожалуйста, введите 'y' или 'n'.")
	}
}

func GenerateReport(g IReportGenerator, logger *Logger) error {
	logger.Log("--- начало генерации отчета ---")

	commonData, err := loadBaseData(logger)
	if err != nil {
		return err
	}

	header, err := g.generateHeader()
	if err != nil {
		return err
	}

	formattedData, err := g.formatData(commonData)
	if err != nil {
		return err
	}

	optionalData := ""
	if g.shouldAddOptionalData() {
		logger.Log("добавление опциональных данных...")
		optionalData, err = g.addOptionalData()
		if err != nil {
			return err
		}
	}

	footer, err := g.generateFooter()
	if err != nil {
		return err
	}

	finalReport := fmt.Sprintf("%s\n%s\n%s\n%s", header, formattedData, optionalData, footer)
	logger.Log("отчет собран.")
	fmt.Println("--- содержимое отчета ---")
	fmt.Println(finalReport)
	fmt.Println("------------------------")

	if g.shouldSave() {
		logger.Log("сохранение отчета...")
		if err := g.save(finalReport); err != nil {
			return err
		}
	} else {
		logger.Log("шаг сохранения пропущен.")
	}

	if g.shouldSendEmail() {
		logger.Log("отправка отчета по email...")
		if err := g.sendEmail(finalReport); err != nil {
			return err
		}
	} else {
		logger.Log("отправка по email пропущена.")
	}

	logger.Log("--- генерация отчета завершена ---")
	return nil
}

type PdfReport struct{}

func (p *PdfReport) generateHeader() (string, error) {
	return "[pdf] заголовок с логотипом компании", nil
}
func (p *PdfReport) formatData(data string) (string, error) {
	return fmt.Sprintf("[pdf] форматирование данных: %s", data), nil
}
func (p *PdfReport) generateFooter() (string, error) {
	return "[pdf] подвал с номерами страниц", nil
}
func (p *PdfReport) shouldAddOptionalData() bool { return true }
func (p *PdfReport) addOptionalData() (string, error) {
	return "[pdf] добавлены метаданные автора", nil
}
func (p *PdfReport) shouldSave() bool      { return true }
func (p *PdfReport) save(data string) error {
	fmt.Println("отчет сохранен как 'report.pdf'")
	return nil
}
func (p *PdfReport) shouldSendEmail() bool { return false }
func (p *PdfReport) sendEmail(data string) error { return nil }

type ExcelReport struct{}

func (e *ExcelReport) generateHeader() (string, error) {
	return "[excel] заголовок (a1: 'название', b1: 'дата')", nil
}
func (e *ExcelReport) formatData(data string) (string, error) {
	return fmt.Sprintf("[excel] данные в ячейках: %s", data), nil
}
func (e *ExcelReport) generateFooter() (string, error) {
	return "[excel] подвал с итоговой суммой", nil
}
func (e *ExcelReport) shouldAddOptionalData() bool { return false }
func (e *ExcelReport) addOptionalData() (string, error) { return "", nil }
func (e *ExcelReport) shouldSave() bool {
	return askUser("сохранить excel отчет?")
}
func (e *ExcelReport) save(data string) error {
	fmt.Println("отчет сохранен как 'report.xlsx'")
	return nil
}
func (e *ExcelReport) shouldSendEmail() bool { return false }
func (e *ExcelReport) sendEmail(data string) error { return nil }

type HtmlReport struct{}

func (h *HtmlReport) generateHeader() (string, error) {
	return "<html><head><title>отчет</title></head><body>\n<h1>заголовок отчета</h1>", nil
}
func (h *HtmlReport) formatData(data string) (string, error) {
	return fmt.Sprintf("<div><p>данные: %s</p></div>", data), nil
}
func (h *HtmlReport) generateFooter() (string, error) {
	return "<footer>(c) 2025 компания</footer>\n</body></html>", nil
}
func (h *HtmlReport) shouldAddOptionalData() bool { return false }
func (h *HtmlReport) addOptionalData() (string, error) { return "", nil }
func (h *HtmlReport) shouldSave() bool { return false }
func (h *HtmlReport) save(data string) error {
	return nil
}
func (h *HtmlReport) sendEmail(data string) error {
	fmt.Printf("html отчет отправлен на 'admin@example.com'\n")
	return nil
}
func (h *HtmlReport) shouldSendEmail() bool {
	return askUser("отправить html отчет по почте?")
}

type CsvReport struct{}

func (c *CsvReport) generateHeader() (string, error) {
	return "id;name;value", nil
}
func (c *CsvReport) formatData(data string) (string, error) {
	return fmt.Sprintf("1;%s;100\n2;%s;200", data, data), nil
}
func (c *CsvReport) generateFooter() (string, error) {
	return "", nil
}
func (c *CsvReport) shouldAddOptionalData() bool { return false }
func (c *CsvReport) addOptionalData() (string, error) { return "", nil }
func (c *CsvReport) shouldSave() bool      { return true }
func (c *CsvReport) save(data string) error {
	fmt.Println("отчет сохранен как 'report.csv'")
	return nil
}
func (c *CsvReport) shouldSendEmail() bool { return false }
func (c* CsvReport) sendEmail(data string) error { return nil }

func main() {
	logger := &Logger{}

	fmt.Println("\n===== генерация pdf отчета =====")
	pdfReport := &PdfReport{}
	if err := GenerateReport(pdfReport, logger); err != nil {
		log.Fatalf("ошибка pdf: %v", err)
	}

	fmt.Println("\n===== генерация excel отчета =====")
	excelReport := &ExcelReport{}
	if err := GenerateReport(excelReport, logger); err != nil {
		log.Fatalf("ошибка excel: %v", err)
	}

	fmt.Println("\n===== генерация html отчета =====")
	htmlReport := &HtmlReport{}
	if err := GenerateReport(htmlReport, logger); err != nil {
		log.Fatalf("ошибка html: %v", err)
	}

	fmt.Println("\n===== генерация csv отчета (расширение) =====")
	csvReport := &CsvReport{}
	if err := GenerateReport(csvReport, logger); err != nil {
		log.Fatalf("ошибка csv: %v", err)
	}
}