# 06 — Strategy + Observer

| Файл | Что демонстрирует |
|---|---|
| `lab/main.go` | **Strategy**: `IShippingStrategy` — расчёт стоимости доставки (Standard/Express/...), выбор способа через интерактивный ввод |
| `prac/main.go` | **Strategy**: `ICostCalculationStrategy` — расчёт стоимости авиабилета в зависимости от класса (`economy`/`business`) и типа пассажира |
| `hometask/main.go` | **Strategy**: `PaymentStrategy` (карта/PayPal/крипта), выбор через `PaymentContext`. **Observer**: `CurrencyExchange` уведомляет `Bank`/`Investor`/`NewsAgency` при изменении курса валют |

```bash
go run "./06 module/hometask/main.go"
```
