package libraries

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func NewValidator() *Validator {
	translato := en.New()
	uni := ut.New(translato, translato)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		Validate: validate,
		Trans:    trans,
	}
}

func (v *Validator) Struct(s interface{}) interface{} {
	errors := make(map[string]string)

	err := v.Validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.StructField()] = e.Translate(v.Trans)
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}
