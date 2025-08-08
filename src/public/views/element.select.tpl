{{define "element.select"}}
<select name="{{.Name}}" id="{{.ID}}">
{{range .Options}}
    <option value="{{.Value}}">{{.Name}}</option>
{{end}}
</select>
{{end}}