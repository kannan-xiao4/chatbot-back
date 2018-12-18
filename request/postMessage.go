package request

type PostMessage struct {
	Message  string `json:"message"`
	UserName string `json:"name"`
	Image    string `json:"image"`
}
