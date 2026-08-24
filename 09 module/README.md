# 09 — Composite, Facade

| Файл | Что демонстрирует |
|---|---|
| `hometask/composite/main.go` | Дерево файловой системы: `File` и `Directory` реализуют общий `FileSystemComponent` (`Display`/`GetSize`/`Name`) — каталог рекурсивно суммирует размер вложенных файлов и папок |
| `hometask/facade/main.go` | `HomeTheaterFacade` прячет за собой `TV`, `AudioSystem`, `DVDPlayer`, `GameConsole` — `WatchMovie()`/`PlayGame()`/`PlayMusic()`/`Shutdown()` включают/выключают нужные устройства одним вызовом вместо ручного управления каждым |

```bash
go run "./09 module/hometask/composite/main.go"
go run "./09 module/hometask/facade/main.go"
```
