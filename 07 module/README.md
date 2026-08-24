# 07 — Command, Mediator, Template Method

Три поведенческих паттерна, у каждого отдельная папка с полным циклом lab → prac → hometask.

## Command (`Command_Pattern/`)

Инкапсуляция действия в объект с `Execute()`/`Undo()` — умный дом (Light, Television, AirConditioner).
- `prac` — добавляет `NoCommand` (null object) и `container/list` для истории команд (отмена по стеку)
- `hometask` — та же идея на упрощённом наборе устройств

## Mediator (`Mediator_Pattern/`)

Центральный посредник вместо прямых связей между объектами — чат-система.
- `lab`/`hometask` — регистрация пользователей, публичные и приватные сообщения, группы
- `prac` — расширенная версия: каналы, бан пользователей, admin-роль, кросс-канальные сообщения

## Template Method (`Template_Method/`)

Скелет алгоритма в базовой функции, шаги — в реализации интерфейса.
- `lab`/`hometask` — классический пример приготовления напитка (`boilWater` → `brew` → `pourInCup` → опционально `addCondiments`)
- `prac` — генератор отчётов (`generateHeader` → `formatData` → `generateFooter`) с логированием каждого шага

```bash
go run "./07 module/lab/Command_Pattern/main.go"
go run "./07 module/lab/Mediator_Pattern/main.go"
go run "./07 module/lab/Template_Method/main.go"
```
