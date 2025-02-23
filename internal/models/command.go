package models

type Command struct {
	Command     string   `json:"command"`
	Description string   `json:"description"`
	Usage       string   `json:"usage"`
	Examples    []string `json:"examples"`
}
