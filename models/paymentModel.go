package models

type PaymentModel struct {
	Amount        float64
	PaymentDate   string
	PaymentMethod string
	Status        string
	BookingId     string
}
