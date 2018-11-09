package root

// NewAthelete structure for new athelete
type NewAthelete struct {
	FriendyURL  string `json:"friendly-url"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date-of-birth"` // dd-MM-yyyy
}

// Athelete structure for  athelete
type Athelete struct {
	HRef         string `json:"href"`
	FriendyURL   string `json:"friendly-url"`
	Name         string `json:"name"`
	DateOfBirth  string `json:"date-of-birth"` // dd-MM-yyyy
	CreationDate string `json:"creation-date"` //date and time of creation
}
