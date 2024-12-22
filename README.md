## Название
HTTP-калькулятор

Этот проект представляет собой HTTP калькулятор, который поддерживает математические операции через HTTP-запросы.

## Возможности
- Сложение
- Вычитание
- Умножение
- Деление
- Вычисление выражений в скобках

## Технологии
- GO
- JSON

Для начала работы с калькулятором необходимо прописать команду, для скачивания всех файлов с репозитория.
go get https://github.com/LootNex/Http_Calc

## Использование

Для запуска сервера необходимо прописать команду 

$env:PORT="9090"; go run cmd/main.go  В таком случае сервер запустить на порту "9090", либо можете указать другой. 

В случае запуска go run cmd/main.go сервер запуститься на порту "8080", который указан по умолчанию.

Для проверки Статуса кода 200 введите команду:
через терминал:

curl.exe --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'

через Postman:

{
    "expression": "2+2"
}


Для проверки Статуса кода 422 введите команду:
через терминал:

curl.exe --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*(2"}'

через Postman:

{
    "expression": "2+(2"
}


Для проверки Статуса кода 500 введите команду:
через терминал:

curl.exe --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"'

через Postman:

{
    "expression": "2+2"


