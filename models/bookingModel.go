package models

type BookingModel struct {
	CheckInDate   string `json:"checkInDate"`
	CheckOutDate  string `json:"checkOutDate"`
	PaymentStatus string `json:"paymentStatus"`
	UserId        string `json:"userId"`
	RoomId        string `json:"roomId"`
	HotelId       string `json:"hotelId"`
	CustomerId    string `json:"customerId"`
}
