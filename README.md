# go-currency-tool

go-currency-tool — небольшая утилита для получения курсов валют с сайта NationalBank.kz, сохранения данных в папке output/ и визуализации CSV-файлов в PNG.

## Что делает проект

- main.go / go-currency-tool.exe: загружает курсы валют, сохраняет их в output/YYYY-MM-DD.csv и output/YYYY-MM-DD.txt.
- Поддерживает вывод курсов в консоль, получение данных за заданную дату, конвертацию валюты в тенге и вызов Python-скрипта для визуализации.
- isualize.py: строит столбчатый график курсов из CSV и сохраняет результат в папке charts/.

## Требования

- Go 1.20+ (или подходящая версия для сборки/запуска main.go)
- Python 3.8+ (для isualize.py)
- Python-библиотеки:
  - pandas
  - seaborn
  - matplotlib

## Установка зависимостей

В PowerShell выполните:

`powershell
py -m pip install -r requirements.txt
`

## Запуск проекта

### 1. Сборка и запуск Go-инструмента

Собрать и запустить:

`powershell
go build -o go-currency-tool.exe .
.\go-currency-tool.exe -p
`

Или запустить напрямую через go run:

`powershell
go run main.go
`

### 2. Примеры использования

- Показать все курсы валют за текущую дату:

`powershell
.\go-currency-tool.exe -p
`

- Получить курсы по дате:

`powershell
.\go-currency-tool.exe -p -d 31.03.2026
`

- Визуализировать CSV-файл:

`powershell
.\go-currency-tool.exe -v .\output\31-03-2026.csv
`

- Конвертировать валюту в тенге:

`powershell
.\go-currency-tool.exe -c USD 100
`

### 3. Запуск визуализации напрямую

`powershell
py .\visualize.py .\output\31-03-2026.csv
`

## Структура проекта

- main.go — основной Go-код для получения и сохранения курсов
- isualize.py — Python-скрипт для создания графиков
- equirements.txt — список Python-зависимостей
- output/ — папка для CSV и TXT файлов
- charts/ — папка для PNG-графиков

