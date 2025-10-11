package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	default:
		return "unknown"
	}
}

type Logger struct {
	LogLevel LogLevel
	filePath string
	mutex    sync.Mutex
}

var (
	instance *Logger
	once     sync.Once
)

func GetLoggerInstance() *Logger {
	once.Do(func() {
		instance = &Logger{
			LogLevel: INFO,
			filePath: "log.txt",
		}
	})
	return instance
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.LogLevel = level
}

func (l *Logger) SetLogFilePath(path string) {
	l.filePath = path
}

func (l *Logger) Log(level LogLevel, message string) {
	if level < l.LogLevel {
		return
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()

	file, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла логов: %v\n", err)
		return
	}
	defer file.Close()

	logLine := fmt.Sprintf("[%s] %s: %s\n", time.Now().Format(time.RFC3339), level.String(), message)
	file.WriteString(logLine)
}

func (l *Logger) ReadLogs() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	data, err := os.ReadFile(l.filePath)
	if err != nil {
		fmt.Println("Ошибка при чтении логов:", err)
		return
	}

	fmt.Println("Содержимое лог-файла:")
	fmt.Println(string(data))
}

func logFromGoroutine(name string, level LogLevel, wg *sync.WaitGroup) {
	defer wg.Done()
	logger := GetLoggerInstance()

	for i := 1; i <= 3; i++ {
		logger.Log(level, fmt.Sprintf("%s — сообщение #%d", name, i))
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	logger := GetLoggerInstance()

	// Настраиваем логгер
	logger.SetLogLevel(WARNING) // INFO < WARNING < ERROR
	logger.SetLogFilePath("logs.txt")

	// Создаем несколько горутин
	var wg sync.WaitGroup
	wg.Add(3)

	go logFromGoroutine("INFO Thread", INFO, &wg)
	go logFromGoroutine("WARNING Thread", WARNING, &wg)
	go logFromGoroutine("ERROR Thread", ERROR, &wg)

	wg.Wait()

	fmt.Printf("\n--- Логирование завершено ---\n")

	// Читаем лог
	logger.ReadLogs()
}
