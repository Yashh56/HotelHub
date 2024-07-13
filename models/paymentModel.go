package models

type PaymentModel struct {
	Amount        float64 `json:"amount"`
	PaymentDate   string  `json:"paymentDate"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	BookingId     string
}
