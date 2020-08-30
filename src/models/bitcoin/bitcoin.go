package bitcoinModel

import "time"

type Bitcoin struct {
	Id            int       `json:"id"`
	Quantity      int       `json:"quantity"`
	Total         float64   `json:"total"`
	Date          time.Time `json:"date"`
	Type          string    `json:"type"`
	Name          string    `json:"name"`
	Person_id     int       `json:"person_id"`
	ValidatedDate string    `json:"validatedDate"`
}

type BitcoinDateReport struct {
	Id        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Date      string  `json:"date"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Person_id int     `json:"person_id"`
}

type BitCoinReportOne struct {
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Person_id int     `json:"person_id"`
	Name      string  `json:"name"`
}

type BitCoinReportOneArray struct {
	Person_id   int                `json:"person_id"`
	Person_name string             `json:"person_name"`
	Bitcoins    []BitCoinReportOne `json:"bitcoins"`
}
