package httprequestpayloadvalidation

import (
	"github.com/go-playground/validator/v10"
)

// Validate
type UserForValidate struct {
	Name  string `json:"name" validate:"required"`                                 //tidak boleh kosong.
	Email string `json:"email" xml:"email" form:"email" validate:"required,email"` // tidak boleh kosong.dan harus format email
	Age   int    `json:"age" xml:"age" form:"age" validate:"gte=0,lte=80"`         //tidak harus di-isi; namun jika ada isinya, maka harus berupa numerik dalam kisaran angka 0 hingga 80.
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
