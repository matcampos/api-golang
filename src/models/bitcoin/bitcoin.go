package bitcoinModel

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

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

func (b Bitcoin) ValidateBitcoinStruct() error {
	return validation.ValidateStruct(&b,
		// Quantity cannot be empty.
		validation.Field(&b.Quantity, validation.Required),
		// Total cannot be empty.
		validation.Field(&b.Total, validation.Required),
		// Type cannot be empty.
		validation.Field(&b.Type, validation.Required),
		// Person_id cannot be empty.
		validation.Field(&b.Person_id, validation.Required),
	)
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
