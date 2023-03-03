package validator

import ut "github.com/go-playground/universal-translator"

// TranslationFunc is the function type used to register or override
// custom translations
type TranslationFunc func(ut ut.Translator, fe FieldError) string

// RegisterTranslationsFunc allows for registering of translations
// for a 'ut.Translator' for use within the 'TranslationFunc'
type RegisterTranslationsFunc func(ut ut.Translator) error

// TranslateValidate is a validator interface to use in translator registration
type TranslateValidate interface {
	RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error)
}

// RegisterDefaultTranslations is the function type that translators should have
type RegisterDefaultTranslations func(v TranslateValidate, trans ut.Translator) (err error)

// RegisterTranslatorDefaultTranslations registers translations in the Translator without needing
// a validator instance
func RegisterTranslatorDefaultTranslations(trans ut.Translator, f RegisterDefaultTranslations) error {
	return f(&translatorTranslateValidate{}, trans)
}

type translatorTranslateValidate struct {
}

func (r translatorTranslateValidate) RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error) {
	return registerFn(trans)
}

// RegisterValidatorDefaultTranslations register translations in the validator without affecting
// the Translator.
// You must have called "RegisterTranslatorDefaultTranslations" on this translator before this call.
func RegisterValidatorDefaultTranslations(v TranslateValidate, trans ut.Translator, f RegisterDefaultTranslations) error {
	return f(&validatorTranslateValidate{v}, trans)
}

type validatorTranslateValidate struct {
	v TranslateValidate
}

func (r validatorTranslateValidate) RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error) {
	return r.v.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return nil // translator registration was already done by RegisterTranslatorDefaultTranslations
	}, translationFn)
}
