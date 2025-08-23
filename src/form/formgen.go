package form

import (
    "bytes"
    "html/template"
    "reflect"
)

type Text string
type Textarea string
type Number int
type Toggle bool
type Select []string
type Checkbox bool

// Field represents a struct field with its tag values.
type Field struct {
    Name        string
    ID          string
    Value       interface{}
    Placeholder string
    ContentType string
    Type        reflect.Type
    Options     []Option
}

// Option represents an option for a select element.
type Option struct {
    Value string
    Label string
}

// FormGenerator is the main type for generating form elements.
type FormGenerator struct {
    tmpl *template.Template
}

// NewFormGenerator creates a new instance of FormGenerator with predefined templates.
func NewFormGenerator() (*FormGenerator, error) {
    tmpl := template.Must(template.New("elements").Parse(`
    {{define "element.input"}}
    <input type="{{.ContentType}}" {{if .Name}}name="{{.Name}}"{{end}} {{if .ID}}id="{{.ID}}"{{end}} value="{{.Value}}" {{if .Placeholder}}placeholder="{{.Placeholder}}"{{end}}>
    {{end}}

    {{define "element.textarea"}}
    <textarea name="{{.Name}}" id="{{.ID}}" placeholder="{{.Placeholder}}">{{.Value}}</textarea>
    {{end}}

    {{define "element.checkbox"}}
    <input type="checkbox" {{if .Name}}name="{{.Name}}"{{end}} {{if .ID}}id="{{.ID}}"{{end}} {{if eq .Value true}}checked{{end}}>
    {{end}}

    {{define "element.select"}}
    <select name="{{.Name}}" id="{{.ID}}">
    {{range .Options}}
    <option value="{{.Value}}" {{if eq $.Value (print .Value)}}selected{{end}}>{{.Label}}</option>
    {{end}}
    </select>
    {{end}}

    {{define "element.toggle"}}
    <label class="switch">
    <input type="checkbox" name="{{.Name}}" id="{{.ID}}" {{if eq .Value true}}checked{{end}}>
    <span class="slider round"></span>
    </label>
    {{end}}
    `))

    return &FormGenerator{
        tmpl: tmpl,
    }, nil
}

// GenerateHTML generates the HTML form elements for a given struct.
func (fg *FormGenerator) GenerateHTML(data interface{}) (string, error) {
    var buf bytes.Buffer

    v := reflect.ValueOf(data)
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    t := v.Type()

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i).Interface()

        fieldInfo := Field{
            Name:        field.Name,
            ID:          field.Tag.Get("id"),
            Value:       value,
            Placeholder: field.Tag.Get("placeholder"),
        }

        switch field.Type.Kind() {
        case reflect.String:
            if field.Type == reflect.TypeOf(Text("")) || field.Type == reflect.TypeOf(Textarea("")) {
                fieldInfo.ContentType = "text"
                if field.Type == reflect.TypeOf(Textarea("")) {
                    fieldInfo.ContentType = "textarea"
                }
                err := fg.tmpl.ExecuteTemplate(&buf, fmtTemplate(fieldInfo.ContentType), fieldInfo)
                if err != nil {
                    return "", err
                }
            }
        case reflect.Int:
            if field.Type == reflect.TypeOf(Number(0)) {
                fieldInfo.ContentType = "number"
                err := fg.tmpl.ExecuteTemplate(&buf, "element.input", fieldInfo)
                if err != nil {
                    return "", err
                }
            }
        case reflect.Bool:
            if field.Type == reflect.TypeOf(Checkbox(false)) || field.Type == reflect.TypeOf(Toggle(false)) {
                fieldInfo.ContentType = "checkbox"
                if field.Type == reflect.TypeOf(Toggle(false)) {
                    err := fg.tmpl.ExecuteTemplate(&buf, "element.toggle", fieldInfo)
                    if err != nil {
                        return "", err
                    }
                    continue
                }
                err := fg.tmpl.ExecuteTemplate(&buf, "element.checkbox", fieldInfo)
                if err != nil {
                    return "", err
                }
            }
        case reflect.Slice, reflect.Array:
            if field.Type == reflect.TypeOf(Select([]string{})) {
                fieldInfo.Options = getOptionsFromField(value)
                err := fg.tmpl.ExecuteTemplate(&buf, "element.select", fieldInfo)
                if err != nil {
                    return "", err
                }
            }
        default:
            // Handle other types if necessary
            continue
        }
    }

    return buf.String(), nil
}

// getOptionsFromField extracts options from a slice or array field.
func getOptionsFromField(value interface{}) []Option {
    var options []Option

    v := reflect.ValueOf(value)
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    for i := 0; i < v.Len(); i++ {
        elem := v.Index(i)
        options = append(options, Option{
            Value: elem.Interface().(string),
            Label: elem.Interface().(string),
        })
    }

    return options
}

// fmtTemplate formats the template name based on content type.
func fmtTemplate(contentType string) string {
    if contentType == "textarea" {
        return "element.textarea"
    }
    return "element.input"
}