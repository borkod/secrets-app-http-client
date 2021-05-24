package cmd

type request struct {
	url    string `json:"url"`
	action string `json:"action"`
	data   string `json:"data"`
	id     string `json:"id"`
}

type secretData struct {
	Data string `json:"data"`
}

type secretID struct {
	Id string `json:"id"`
}
