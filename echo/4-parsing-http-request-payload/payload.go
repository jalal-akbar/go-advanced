package parsinghttprequestpayload

// Payload
type UserForPayload struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
}
