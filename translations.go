package validator

import ut "github.com/go-playground/universal-translator"

// TranslationFunc is the function type used to register or override
// custom translations
type TranslationFunc func(ut ut.Translator, fe FieldError) string

// RegisterTranslationsFunc allows for registering of translations
// for a 'ut.Translator' for use within the 'TranslationFunc'
type RegisterTranslationsFunc func(ut ut.Translator) error

type TranslateValidate interface {
	RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error)
}

type RegisterDefaultTranslations func(v TranslateValidate, trans ut.Translator) (err error)

func RegisterTranslatorDefaultTranslations(trans ut.Translator, f RegisterDefaultTranslations) error {
	return f(&registerOnlyTranslateValidate{}, trans)
}

type registerOnlyTranslateValidate struct {
}

func (r registerOnlyTranslateValidate) RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error) {
	return registerFn(trans)
}

func RegisterValidatorDefaultTranslations(v TranslateValidate, trans ut.Translator, f RegisterDefaultTranslations) error {
	return f(&validatorTranslateValidate{v}, trans)
}

type validatorTranslateValidate struct {
	v TranslateValidate
}

func (r validatorTranslateValidate) RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) (err error) {
	return r.v.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return nil
	}, translationFn)
	return nil
}
