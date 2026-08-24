# 03 — SOLID

Все четыре принципа SOLID (кроме LSP) на прикладных примерах.

| Принцип | Пример в `lab/main.go` |
|---|---|
| **SRP** | `Invoice` отдельно от `InvoiceCalculator` и `InvoiceRepository` — расчёт, хранение и данные не смешаны |
| **OCP** | `DiscountStrategy` — новый тип скидки (`RegularCustomer`/`SilverCustomer`/`GoldCustomer`) добавляется без изменения `DiscountCalculator` |
| **ISP** | `Workable`/`Eatable`/`Sleepable` — `RobotWorker` реализует только `Work()`, не обязан "уметь есть и спать" |
| **DIP** | `Notification` зависит от интерфейса `Sender`, а не от конкретного `EmailService`/`SmsService` |

`prac/main.go` — интернет-магазин: `Order` собирает оплату (`IPayment`), доставку (`IDelivery`), уведомления (`INotification`) и скидки (`IDiscount`) через интерфейсы — расширяемо без правок `Order`.

`hometask/main.go` — те же 4 принципа на других примерах: расчёт зарплаты сотрудников (OCP), принтер/сканер/факс (ISP), сервис уведомлений (DIP).

```bash
go run "./03 module/lab/main.go"
```
