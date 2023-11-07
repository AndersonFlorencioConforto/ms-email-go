package main

import (
	"github.com/go-playground/validator/v10"
	"ms-email/internal/domain/campaign"
)

func main() {

	contacts := []campaign.Contact{
		{
			Email: "",
		},
	}
	campaign := campaign.Campaign{
		Contacts: contacts,
	}

	var validate = validator.New()

	err := validate.Struct(campaign)
	if err == nil {
		println("No errors, campaign is valid")
	} else {
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			println(value.StructField() + " " + value.Tag() + " " + value.Param())
		}
	}
}
