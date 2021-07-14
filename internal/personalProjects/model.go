package personalProjects

type PersonalProjects struct {
	Id          int    `json:"Id,omitempty"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Url         string `json:"Url"`
}
