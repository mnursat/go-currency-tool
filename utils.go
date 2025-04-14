package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// PrintAll универсально выводит элементы с цветом
func printAll[T any](items []T) {
	for _, item := range items {
		switch v := any(item).(type) {
		case Item:
			fmt.Printf("Валюта: %s | Курс: %s | Кол-во: %d\n", v.Title, v.Description, v.Quant)
		case RateItem:
			fmt.Printf("Валюта: %s | Курс: %s | Кол-во: %d\n", v.Title, v.Description, v.Quant)
		}
	}
}

// writeToTxt универсально сохраняет данные в .txt
func writeToTxt[T any](path string, items []T) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, item := range items {
		switch v := any(item).(type) {
		case Item:
			fmt.Fprintf(f, "Валюта: %s | Курс: %s | Кол-во: %d\n", v.Title, v.Description, v.Quant)
		case RateItem:
			fmt.Fprintf(f, "Валюта: %s | Курс: %s | Кол-во: %d\n", v.Title, v.Description, v.Quant)
		}
	}
	return nil
}

// WriteToCsv универсально сохраняет данные в .csv
func writeToCsv[T any](path string, items []T) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Заголовки
	w.Write([]string{"Валюта", "Курс", "Кол-во"})

	for _, item := range items {
		switch v := any(item).(type) {
		case Item:
			w.Write([]string{v.Title, v.Description, strconv.Itoa(v.Quant)})
		case RateItem:
			w.Write([]string{v.Title, v.Description, strconv.Itoa(v.Quant)})
		default:
			// Можешь логировать/игнорировать неизвестный тип
			continue
		}
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
