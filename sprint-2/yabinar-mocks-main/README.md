# Webinar #2

# Моки

## Написание моков вручную

См. `internal/core/application/application_mock.go`

## Моки без `go:generate`

Установите gomock локально:

```bash
go install go.uber.org/mock/mockgen@latest
```

Моки можно сгенерировать, передавая путь к файлу с интерфейсами:

```bash
mockgen -source=internal/core/services/repo.go \
    -destination=internal/core/services/internal/mocks/repo_mock.gen.go \
    -package=mocks
```

А можно сгенерировать, передавая import path пакета и список интерфейсов для мокирования:

```bash
mockgen -destination=internal/core/services/internal/mocks-reflect/repo_mock.gen.go \
    -package=mocks \
    mocks/internal/core/services Store,Foobar
```

Разница между source-mode и reflect-mode:

1. source-mode работает с неэкспортированными интерфейсами
2. reflect-mode позволяет выбирать интерфейсы для мокирования
3. reflect-mode работает в случае сложной организации репозитория (например, когда модули не используются)

См. обсуждение: https://github.com/golang/mock/issues/406

См. пример использования моков в тестах в `internal/core/services/repo_test.go`

## Моки с `go:generate`

Для генерации моков можно использовать `go:generate` (см. https://go.dev/blog/generate).

См. пример в `internal/core/services/repo_test.go`.

Для генерации кода выполните:

```bash
go generate ./...
```

в корне проекта.

## Mockery

Другой популярный пакет для создания моков (см. сравнение их и других пакетов для мокирования здесь: https://gist.github.com/maratori/8772fe158ff705ca543a0620863977c2).

Установите `mockery` при помощи:

```bash
go install github.com/vektra/mockery/v2@v2.43.1
```

Конфигурация `mockery` находится в `.mockery.yaml`.

Сгенерируйте код:

```bash
mockery
```

Тесты с использованием `mockery` лежат в `internal/core/services/repo_mockery_test.go`.

# Makefile

Дока по Makefile лежит здесь: https://www.gnu.org/software/make/manual/make.html

См. Makefile в корне проекта

# Линтеры

## `golangci-lint`

См. https://golangci-lint.run/

Для запуска линтеров выполните:

```bash
make lint
```

Для того, чтобы запустить `golangci-lint` с файлом конфигурации выполните:

```bash
make lint GOLANGCI_LINT_CONFIG_FILE=lint.yml
```
