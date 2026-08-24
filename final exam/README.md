# Final Exam — ISP + Adapter

Демонстрация нарушения **Interface Segregation Principle** и его исправления через **Adapter**.

```go
type IBadWorker interface {
	Work()
	Eat()
	Sleep()
}
```

`Robot` вынужден реализовывать `Eat()`/`Sleep()`, хотя роботы не едят и не спят — эти методы паникуют (`panic("Роботы не едят")`). Классическая проблема "толстого" интерфейса.

**Решение:** узкий интерфейс `Worker { Work() }` + `WorkerAdapter`, который оборачивает `IBadWorker` и наружу отдаёт только `Work()`. И `Human`, и `Robot` используются единообразно через `Worker`, не будучи вынужденными знать о методах друг друга.

```bash
go run "./final exam/main.go"
```
