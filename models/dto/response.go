package dto
type ResponseDto struct {
	Data any `json:"data,omitempty"`
	Code string `json:"code"`
	Message string `json:"message"`
	Error string `json:"error,omitempty"`
	Errors *[]string `json:"errors,omitempty"`
	HttpResponseCode int `json:"-"`
}