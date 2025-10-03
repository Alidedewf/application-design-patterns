package main

import (
	"fmt"
	"sort"
)

// DRY
var connectionString = "Server=myServer;Database=myDb;User Id=myUser;Password=myPass;"

func log(level, message string) {
	fmt.Printf("%s: %s\n", level, message)
}

type DatabaseService struct{}

func (d *DatabaseService) Connect() {
	fmt.Println("Connecting with:", connectionString)
}

type LoggingService struct{}

func (l *LoggingService) Log(message string) {
	fmt.Println("Logging with:", connectionString)
	fmt.Println("Message:", message)
}

// KISS
func ProcessNumbers(numbers []int) {
	if len(numbers) == 0 {
		return
	}
	for _, n := range numbers {
		if n > 0 {
			fmt.Println(n)
		}
	}
}

func PrintPositiveNumbers(numbers []int) {
	var positives []int
	for _, n := range numbers {
		if n > 0 {
			positives = append(positives, n)
		}
	}
	sort.Ints(positives)
	for _, n := range positives {
		fmt.Println(n)
	}
}

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// YAGNI
type user struct {
	Name  string
	Email string
	Role  string
}

func ReadFile(filePath string) string {
	fmt.Println("Reading file:", filePath)
	return "file content"
}

type ReportGenerator struct{}

func (r *ReportGenerator) GeneratePDF() {
	fmt.Println("Generating PDF report")
}

func main() {
	// DRY
	log("INFO", "App started")
	db := DatabaseService{}
	db.Connect()
	logger := LoggingService{}
	logger.Log("User logged in")

	// KISS
	numbers := []int{5, -1, 0, 8, 3}
	ProcessNumbers(numbers)
	PrintPositiveNumbers(numbers)
	fmt.Println("Divide result:", Divide(10, 0))

	// YAGNI
	user := user{Name: "Alihan", Email: "ali@example.com", Role: "User"}
	fmt.Println("User:", user.Name)

	content := ReadFile("file.txt")
	fmt.Println("File content:", content)

	report := ReportGenerator{}
	report.GeneratePDF()
}
