# 01 — ООП: наследование и композиция

Базовые механизмы Go для переиспользования кода через встраивание структур (embedding).

| Файл | Что демонстрирует |
|---|---|
| `lab/main.go` | `Vehicle` → `Car` — встраивание базовой структуры |
| `prac/main.go` | `Employee` → `Worker` / `Manager` — общее поле через embedding, разная реализация `CalculateSalary()` |
| `hometask/main.go` | Мини-система библиотеки: `Library` управляет `Book`/`Reader`, выдача и возврат книг с учётом количества экземпляров |

```bash
go run ./01\ module/hometask/main.go
```
