package model

import (
	"reflect"
	"sync"

	// import validator v9
	// validator "gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	ja_translations "gopkg.in/go-playground/validator.v9/translations/ja"
)

// Translator translates validation messages
var Translator ut.Translator

// StructValidator defines validations for models.
type StructValidator struct {
	once     sync.Once
	validate *validator.Validate
	uni      *ut.UniversalTranslator
}

// ValidateStruct validates struct with tags.
func (v *StructValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

// Engine returns validate engine.
func (v *StructValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}
func (v *StructValidator) lazyinit() {
	v.once.Do(func() {
		ja := ja.New()
		en := en.New() // fallback
		v.uni = ut.New(en, ja)

		trans, _ := v.uni.GetTranslator("ja")
		Translator = trans
		v.validate = validator.New()

		ja_translations.RegisterDefaultTranslations(v.validate, trans)

		v.validate.SetTagName("binding")
		// add any custom validations etc. here

	})
}
func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
