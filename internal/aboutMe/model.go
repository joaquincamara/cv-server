package aboutMe

type AboutMe struct {
	Id    int    `json:"Id,omitempty"`
	Title string `json:"Title"`
	Info  string `json:"Info"`
}
