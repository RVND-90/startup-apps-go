package dto

import "github.com/RVND-90/startup-apps-go/constants"
type ResponseDto struct {
	Data any `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
	Errors *[]string `json:"errors,omitempty"`
	*constants.RC
}

