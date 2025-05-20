# go_modules_cli
Анализ зависимостей репозитория
Задача: Для указанного репозитория вывести данные о модуле и список зависимостей, которые можно обновить

## Установка

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/xddprog/go_module_cli.git
   cd go_module_cli
   ```

## Использование

Запустите утилиту, указав URL Git-репозитория:

```bash
go run cmd/main.go <git-repo-url>
```

Пример вывода:

```plaintext
Module: github.com/xddprog/real-time-text-editor
Go Version: 1.24.2
Updating github.com/gorilla/mux from v1.7.4 to v1.8.1
Updating github.com/gorilla/websocket from v1.4.2 to v1.5.3
Updating github.com/jackc/pgtype from v1.14.0 to v1.14.4
Updating github.com/jackc/pgx/v4 from v4.18.2 to v4.18.3
Updating github.com/jackc/pgx/v5 from v5.7.4 to v5.7.5
```

Если обновлений нет:

```plaintext
Module: github.com/xddprog/real-time-text-editor
Go Version: 1.18
Dependencies:
No updates available
```

Если нет данных о модуле:
```
Module: unkown
Go Version: unkown
Dependencies:
...
```

Для локального пути:

```bash
go run cmd/go_modules_cli/main.go /path/to/local/repo
```

## Структура проекта

- `cmd/go_modules_cli/main.go`: Точка входа CLI, оркестрация модулей.
- `internal/repo`: Загрузка репозитория (Git-клонирование или локальный путь).
- `internal/modfile`: Парсинг файла `go.mod` для получения имени модуля и версии Go.
- `internal/dependencies`: Проверка обновлений зависимостей через `go list`.
