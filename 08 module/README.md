# 08 — Adapter, Decorator

## Adapter

Приведение несовместимых интерфейсов внешних сервисов к одному общему.

| Файл | Пример |
|---|---|
| `lab/adapter/main.go` | Внутренний `IPaymentProcessor` vs внешняя `ExternalPaymentSystemA` с другой сигнатурой методов |
| `prac/logistics/main.go` | Три разных внешних логистических сервиса (A/B/C, у каждого свои методы) под один интерфейс `IInternalDeliveryService` |
| `hometask/Adapter Pattern/main.go` | `PayPalPaymentProcessor` vs `StripePaymentService` (`MakeTransaction` вместо `ProcessPayment`) |

## Decorator

Динамическое добавление поведения без изменения исходного класса.

| Файл | Пример |
|---|---|
| `lab/decorator/main.go` | Кофе + `MilkDecorator` (и другие добавки), каждый слой оборачивает `IBeverage` |
| `prac/reports/main.go` | Цепочка декораторов отчёта: `DateFilterDecorator` → `SortingDecorator` → `CsvExportDecorator` |
| `hometask/Decorator Pattern/main.go` | Напитки (`Americano`/`Tea`) + декораторы добавок, аналогично lab |

```bash
go run "./08 module/prac/logistics/main.go"
go run "./08 module/prac/reports/main.go"
```
