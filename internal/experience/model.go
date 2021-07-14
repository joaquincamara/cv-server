package experience

type Experience struct {
	Id       int    `json:"Id,omitempty"`
	Jobtitle string `json:"JobTitle"`
	Date     string `json:"Date"`
}
