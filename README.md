# nyancatsay

Аналог `cowsay` на Go с Nyan Cat. Использует только стандартную библиотеку.

## Запуск

```sh
go run . Привет, мир!
go run . -rainbow Привет, мир!
printf 'Nyan\nNyan' | go run . -rainbow
```

Без аргументов текст читается из stdin. Аргументы объединяются пробелами, переводы строк из stdin сохраняются.

## Сборка

```sh
go build -o nyancatsay .
./nyancatsay -rainbow 'Hello, world!'
```

Флаг `-rainbow` включает ANSI-цвета; без него выводится обычный текст.

## Проверка

```sh
go test ./...
```
