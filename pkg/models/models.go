package models

import (
	"encoding/xml"
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching records found")
)

type Schedule struct {
	XMLName xml.Name `xml:"working_time"`
	Open    string   `xml:"open"`
	Close   string   `xml:"close"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	ID          int      `xml:"id,attr"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	Price       float64  `xml:"price"`
}

type Shop struct {
	XMLName     xml.Name `xml:"shop"`
	ID          int      `xml:"id,attr"`
	Name        string   `xml:"name"`
	URL         string   `xml:"url"`
	WorkingTime Schedule
	Offers      []*Item `xml:"offers>item"`
}
