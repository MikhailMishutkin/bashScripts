package models

type Script struct {
	UUID     int      `json:"UUID,omitempty"`
	Name     string   `json:"name"`
	Commands []string `json:"commands,omitempty"`
	Result   string   `json:"result,omitempty"`
}
