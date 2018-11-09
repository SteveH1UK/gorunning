package root

// NewResource - application new resource created
type NewResource struct {
	HTTPCode int    `json:"http-code"`
	ID       string `json:"id"`
	Href     string `json:"href"`
}
