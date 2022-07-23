package business

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)


func validation(T any) error {
	validate := validator.New()

	err := validate.Struct(T)
	if err != nil {
		t := translateError(err)
		return t
	}

	return nil
}

func translateError(errs error) error {
	var uni *ut.UniversalTranslator

	err := errs.(validator.ValidationErrors)
	
	trans, _ := uni.GetTranslator("en")

	translatedErr := err[0].Translate(trans)

	return errors.New(translatedErr)
}	