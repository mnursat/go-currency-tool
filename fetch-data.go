package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

// Получение всех текущих курсов из https://nationalbank.kz/rss/rates_all.xml
func fetchCurrencyData() ([]RateItem, error) {
	resp, err := http.Get("https://nationalbank.kz/rss/rates_all.xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss struct {
		Channel []struct {
			Items []struct {
				Title       string `xml:"title"`
				Description string `xml:"description"`
				Quant       int    `xml:"quant"`
				Index       string `xml:"index"`
				Change      string `xml:"change"`
			} `xml:"item"`
		} `xml:"channel"`
	}

	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		return nil, err
	}

	var items []RateItem
	for _, item := range rss.Channel[0].Items {
		items = append(items, RateItem{
			Title:       item.Title,
			Description: item.Description,
			Quant:       item.Quant,
			Index:       item.Index,
			Change:      item.Change,
		})
	}
	return items, nil
}

// Получение всех курсов по дате из https://nationalbank.kz/rss/get_rates.cfm?fdate=dd.MM.yyyy
func fetchCurrencyDataByDate(date string) ([]RateItem, error) {
	url := fmt.Sprintf("https://nationalbank.kz/rss/get_rates.cfm?fdate=%s", date)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rates struct {
		Items []RateItem `xml:"item"`
	}

	if err := xml.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, err
	}

	if len(rates.Items) == 0 {
		return nil, fmt.Errorf("данные за дату %s не найдены", date)
	}

	return rates.Items, nil
}

func getExchangeRateByName(currency, date string) (float64, int, error) {
	url := fmt.Sprintf("https://nationalbank.kz/rss/get_rates.cfm?fdate=%s", date)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var rss Rates
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return 0, 0, err
	}

	for _, item := range rss.Items {
		if item.Title == currency {
			rate, err := strconv.ParseFloat(item.Description, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("не удалось преобразовать курс валюты %s", currency)
			}
			return rate, item.Quant, nil
		}
	}
	return 0, 0, fmt.Errorf("курс для валюты %s не найден", currency)
}
