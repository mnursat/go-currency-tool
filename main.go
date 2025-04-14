package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Получаем флаги
	printFlag := flag.Bool("p", false, "Вывыести все данные в консоль")
	visualizePath := flag.String("v", "", "Путь к .csv-файлу для визуализации")
	dateFlag := flag.String("d", "", "Получить курсы валют по дате в формате дд.мм.гггг")
	convertFlag := flag.Bool("c", false, "Конвертация валюты в тенге (например, USD 100)")

	flag.Usage = func() {
		fmt.Println("go-currency-tool — утилита для получения и визуализации курсов валют с сайта NationalBank.kz")
		fmt.Println()
		fmt.Println("Использование:")
		fmt.Println("  ./go-currency-tool [флаги и параметры]")
		fmt.Println()
		fmt.Println("Флаги:")
		fmt.Println("  -h              Показать справку")
		fmt.Println("  -p              Показать все курсы валют в консоли (текущая дата)")
		fmt.Println("  -d ДАТА         Получить курсы по дате (формат: дд.мм.гггг)")
		fmt.Println("  -v              Визуализировать курс валют (текущая дата или с флагом -d)")
		fmt.Println("  -v ПУТЬ         Визуализировать курс валют из указанного CSV-файла")
		fmt.Println("  -c ВАЛЮТА КОЛ-ВО   Конвертировать валюту в тенге, например: -c USD 100")
		fmt.Println()
		fmt.Println("Примеры:")
		fmt.Println("  ./go-currency-tool -p")
		fmt.Println("  ./go-currency-tool -p -d 12.02.2020")
		fmt.Println("  ./go-currency-tool -d 12.02.2020 -v")
		fmt.Println("  ./go-currency-tool -c usd 50")
		fmt.Println("  ./go-currency-tool -v ./output/09-06-2000.csv")
	}

	flag.Parse()

	// Устанавливаем дату
	date := time.Now().Format("02-01-2006")
	if *dateFlag != "" {
		date = strings.ReplaceAll(*dateFlag, ".", "-")
	}

	dir := "output"
	txtPath := filepath.Join(dir, date+".txt")
	csvPath := filepath.Join(dir, date+".csv")
	os.MkdirAll(dir, os.ModePerm)

	// Конвертация валюты
	if *convertFlag {
		if len(flag.Args()) < 2 {
			log.Fatal("Ошибка: для флага -c нужно указать валюту и количество, например, -c USD 100")
		}

		currency := strings.ToUpper(flag.Args()[0])
		amount, err := strconv.Atoi(flag.Args()[1])

		if err != nil {
			log.Fatal("Ошибка при конвертации количества валюты:", err)
		}
		rate, quant, err := getExchangeRateByName(currency, date)
		if err != nil {
			log.Fatalf("Ошибка получения курса валюты для %s: %v", currency, err)
		}
		converted := float64(amount) * rate * float64(quant)
		log.Printf("Конвертация %d %s в тенге: %.2f KZT\n", amount, currency, converted)
		return
	}

	// Печать всех данных
	if *printFlag {
		if *dateFlag != "" {
			items, err := fetchCurrencyDataByDate(*dateFlag)
			if err != nil {
				log.Fatal("Ошибка получения данных:", err)
			}
			printAll(items)
			return
		}

		items, err := fetchCurrencyData()
		if err != nil {
			log.Fatal("Ошибка получения данных:", err)
		}
		printAll(items)
		return
	}

	// Визуализация
	if *visualizePath != "" {
		if !exists(*visualizePath) {
			log.Fatalf("Файл %s не найден", *visualizePath)
		}
		runPythonVisualizer(*visualizePath)
		return
	}

	// Проверка существования файлов
	if exists(txtPath) && exists(csvPath) {
		log.Println("Файлы уже существуют. Программа завершена.")
		return
	}

	// Получение и сохранение данных
	var (
		items []RateItem
		err   error
	)

	if *dateFlag != "" {
		items, err = fetchCurrencyDataByDate(*dateFlag)
	} else {
		items, err = fetchCurrencyData()
	}
	if err != nil {
		log.Fatal("Ошибка получения данных:", err)
	}

	writeToTxt(txtPath, items)
	writeToCsv(csvPath, items)

	log.Println("✅ Готово. Данные сохранены в", dir)
}

func runPythonVisualizer(csvPath string) {
	cmd := exec.Command("python3", "visualize.py", csvPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("📊 Запуск визуализации через Python...")
	if err := cmd.Run(); err != nil {
		log.Fatal("❌ Ошибка при запуске визуализации:", err)
	}
}
