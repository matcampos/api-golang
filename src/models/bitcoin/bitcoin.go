package bitcoinModel

import "time"

type Bitcoin struct {
	Id            int
	Quantity      int
	Total         float64
	Date          time.Time
	Type          string
	Name          string
	Person_id     int
	ValidatedDate string
}

type BitcoinDateReport struct {
	Id        int
	Quantity  int
	Total     float64
	Date      string
	Type      string
	Name      string
	Person_id int
}

type BitCoinReportOne struct {
	Quantity  int
	Total     float64
	Person_id int
	Name      string
}

type BitCoinReportOneArray struct {
	Person_id   int
	Person_name string
	Bitcoins    []BitCoinReportOne
}
