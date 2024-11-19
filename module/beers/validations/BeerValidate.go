package validations


import (
	"beer/module/beers/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateBeerData(beer *models.CreateBeer) error {
	return validate.Struct(beer)
}