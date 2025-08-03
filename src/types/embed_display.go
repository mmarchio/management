package types

type Display struct {
	ContentType string
	DisplayName string
	Form []FormElement
	FormAction string
	FormMethod string
}

type FormElement struct {
	ContentType string
	Name string
	ID string
	Value interface{}
	Checked string
	Placeholder string
	Options []Option
}

type Option struct {
	Name string
	ID string
}