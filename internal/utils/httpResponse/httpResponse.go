package httpResponse

type Success struct {
	Status int `json:"-"`

	Message string `json:"message,omitempty"`
	Item    any    `json:"item,omitempty"`
	Items   any    `json:"items,omitempty"`
	Code    string `json:"code,omitempty"`
}
