package main

type RSS struct {
	XMLName string    `xml:"rss"`
	Channel []Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`       // Название валюты, например "USD"
	PubDate     string `xml:"pubDate"`     // Дата публикации, например "12.04.2025"
	Description string `xml:"description"` // Курс валюты, например "116.5"
	Quant       int    `xml:"quant"`       // Кол-во единиц, например 1
	Index       string `xml:"index"`       // Тип индекса (например, "UP")
	Change      string `xml:"change"`      // Изменение курса, например +1.14
}

type Rates struct {
	XMLName     string     `xml:"rates"`
	Generator   string     `xml:"generator"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Copyright   string     `xml:"copyright"`
	Date        string     `xml:"date"`
	Items       []RateItem `xml:"item"`
}

type RateItem struct {
	FullName    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       int    `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}
