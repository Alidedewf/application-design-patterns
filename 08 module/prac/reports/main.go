package main

import (
	"fmt"
)

type IReport interface {
	Generate() string
}

type SalesReport struct{}
func (r *SalesReport) Generate() string {
	return "Sales Report: [Sale1: $100, Sale2: $200, Sale3: $150]"
}

type UserReport struct{}
func (r *UserReport) Generate() string {
	return "User Report: [User1, User2, User3]"
}

type ReportDecorator struct {
	report IReport
}

func (d *ReportDecorator) Generate() string {
	return d.report.Generate()
}

// Фильтрация по датам
type DateFilterDecorator struct {
	ReportDecorator
	startDate, endDate string
}

func (d *DateFilterDecorator) Generate() string {
	base := d.report.Generate()
	return fmt.Sprintf("%s\n[Filtered by dates: %s - %s]", base, d.startDate, d.endDate)
}

// Сортировка
type SortingDecorator struct {
	ReportDecorator
	criteria string
}

func (s *SortingDecorator) Generate() string {
	base := s.report.Generate()
	return fmt.Sprintf("%s\n[Sorted by: %s]", base, s.criteria)
}

// Экспорт CSV
type CsvExportDecorator struct {
	ReportDecorator
}

func (e *CsvExportDecorator) Generate() string {
	base := e.report.Generate()
	return fmt.Sprintf("%s\n[Exported as: CSV]", base)
}

// Экспорт PDF
type PdfExportDecorator struct {
	ReportDecorator
}

func (e *PdfExportDecorator) Generate() string {
	base := e.report.Generate()
	return fmt.Sprintf("%s\n[Exported as: PDF]", base)
}

func main() {
	var report IReport

	report = &SalesReport{}

	report = &DateFilterDecorator{ReportDecorator{report}, "2025-01-01", "2025-12-31"}
	report = &SortingDecorator{ReportDecorator{report}, "amount"}
	report = &CsvExportDecorator{ReportDecorator{report}}

	fmt.Println(report.Generate())
}