package bitcoinModel

import "time"

type Bitcoin struct {
	Id       int
	Quantity int
	Total    float64
	Date     time.Time
}
