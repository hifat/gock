package uResponse

type SuccessResponse struct {
	Status int `json:"-"`

	Message string `json:"message,omitempty"`
	Item    any    `json:"item,omitempty"`
	Items   any    `json:"items,omitempty"`
	Code    string `json:"code,omitempty"`
}
