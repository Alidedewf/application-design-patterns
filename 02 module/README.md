# 02 — DRY, KISS, YAGNI

Принципы чистого кода на конкретных примерах: устранение дублирования, минимальная сложность, отказ от лишней функциональности.

| Файл | Что демонстрирует |
|---|---|
| `lab/main.go` | Все три принципа: `OrderService` (DRY — общий `calculateTotal`), `Add`/`Singleton` (KISS), `CircleArea`/`SquareArea` (YAGNI — не усложняем там, где не нужно) |
| `prac/main.go` | `UserManager` — CRUD без дублирования логики поиска пользователя |
| `hometask/main.go` | То же трио на примере `DatabaseService`/`LoggingService` (DRY), обработки чисел (KISS), генератора отчётов (YAGNI) |

```bash
go run "./02 module/lab/main.go"
```
