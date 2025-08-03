{{define "element.toggle"}}
<span>{{if .Title}}{{.Title}}{{end}}</span><label class="switch"><input type="checkbox" name="{{if .NamePrefix}}{{.NamePrefix}}{{end}}{{if .Suffix}}{{.Suffix}}{{end}}" id="{{if .IdPrefix}}{{.IdPrefix}}{{end}}{{if .Suffix}}{{.Suffix}}{{end}}"{{if eq .Value true}}checked{{end}}><span class="slider round"></span></label>
{{end}}