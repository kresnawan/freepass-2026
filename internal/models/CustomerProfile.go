package models

type CustomerProfile struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	CanteenPoints int    `json:"canteen_points"`
	Instagram     string `json:"instagram"`
	Bio           string `json:"bio"`
}
