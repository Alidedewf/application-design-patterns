# 04 — Factory Method

Создание объектов через фабрику вместо прямого вызова конструктора — тип создаваемого объекта решается в рантайме.

| Файл | Что демонстрирует |
|---|---|
| `lab/main.go` | `TransportFactory` → `CarFactory`/`MotorcycleFactory`/`PlaneFactory`, каждая создаёт свою реализацию `ITransport` |
| `prac/main.go` | `DocumentCreator` — выбор `ReportCreator`/`ResumeCreator`/`LetterCreator`/`InvoiceCreator` по строке, введённой пользователем |
| `hometask/main.go` | Интерактивный ввод (`bufio.Reader`) — пользователь выбирает тип транспорта (car/motorcycle/truck/bus) и параметры, `VehicleFactory` собирает нужный объект |

```bash
go run "./04 module/hometask/main.go"
```
