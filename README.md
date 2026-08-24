# Application Design Patterns

> Учебная работа — лабораторные, практики и домашние задания университетского курса по паттернам проектирования (Go).

От базовых принципов ООП (DRY, SRP) до классических паттернов GoF — Singleton, Strategy, Command, Mediator, Template Method, Adapter, Decorator, Composite, Facade — каждый модуль на отдельном прикладном примере (доставка, платежи, отчёты, логирование).

## Запуск

```bash
go run "./XX module/lab/main.go"
```
(аналогично для `hometask` и `prac` внутри каждого модуля)

## Модули

У каждого модуля свой README с разбором конкретных файлов — открывается автоматически при переходе в папку на GitHub.


| Модуль | Тема | Примеры |
|---|---|---|
| [01](01%20module) | ООП: наследование и композиция | Vehicle/Car, Employee/Worker, Book/Reader |
| [02](02%20module) | Принцип DRY | Сервис заказов, менеджер пользователей, работа с БД |
| [03](03%20module) | Принцип SRP | Счета-фактуры, обработка платежей, расчёт заказов |
| [04](04%20module) | Интерфейсы и полиморфизм | Транспорт, документы |
| [05](05%20module) | **Singleton** (потокобезопасный логгер, `sync`) | `lab/logger` |
| [06](06%20module) | **Strategy** | Способ доставки, способ оплаты, расчёт стоимости бронирования |
| [07](07%20module) | **Command**, **Mediator**, **Template Method** | Отдельная папка на каждый паттерн (lab / prac / hometask) |
| [08](08%20module) | **Adapter**, **Decorator** | 3 внешних логистических сервиса под один интерфейс; цепочка декораторов отчёта (фильтр по датам → сортировка → экспорт в CSV) |
| [09](09%20module) | **Composite**, **Facade** | Дерево файловой системы, единая точка доступа к отчётам |
| [12](12%20module) | UML-диаграммы паттернов (`.drawio`) | Черновики к лабам/практикам/дз модуля |
| [15](15%20module) | Adapter (повторное закрепление) | — |
| [final exam](final%20exam) | **ISP + Adapter**: демонстрация нарушения Interface Segregation Principle (`IBadWorker` заставляет `Robot` реализовывать `Eat()`/`Sleep()`) и его исправление через `WorkerAdapter` | — |

Модули 05 (Builder) и 16 — незавершённые заготовки (только объявление пакета, без реализации).

## Показательный пример: Interface Segregation + Adapter (final exam)

```go
type IBadWorker interface {
	Work()
	Eat()
	Sleep()
}
// Robot.Eat() и Robot.Sleep() вызывают panic — интерфейс навязывает лишнее.

type Worker interface{ Work() } // узкий интерфейс

type WorkerAdapter struct{ worker IBadWorker }
func (a *WorkerAdapter) Work() { a.worker.Work() }
// Human и Robot работают через общий Worker, не зная о "лишних" методах друг друга.
```

## Пример: цепочка декораторов (08 module/prac/reports)

```go
report := &SalesReport{}
withDates  := &DateFilterDecorator{ReportDecorator{report}, "2024-01-01", "2024-12-31"}
withSort   := &SortingDecorator{ReportDecorator{withDates}, "date"}
withExport := &CsvExportDecorator{ReportDecorator{withSort}}
withExport.Generate() // отчёт → фильтр → сортировка → CSV, каждый слой не знает о соседних
```
