{{define "element.input"}}
<input type="{{.ContentType}}" {{if .Name}}name="{{.Name}}"{{end}} {{if .ID}}id="{{.ID}}"{{end}} value="{{.Value}}" {{if .Placeholder}}placeholder="{{.Placeholder}}"{{end}}>
{{end}}