{{define "form.disposition"}}
<ul>
    <li><input type="text" name="name" id="name" value="{{if .Name}}{{.Name}}{{end}}" placeholder="name"></li>
    <li><input type="number" name="min_duration" id="min_duration" value="{{if .MinDuration}}{{.MinDuration}}{{end}}" placeholder="min duration"></li>
    <li><input type="number" name="max_duration" id="max_duration" value="{{if .MaxDuration}}{{.MaxDuration}}{{end}}" placeholder="max duration"></li>
    <li><input type="number" name="advertisement_duration" id="advertisement_duration" value="{{if .AdvertisementDuration}}{{.AdvertisementDuration}}{{end}}" placeholder="advertisement duration"></li>
    <li>{{if .Entitlements}}{{template "entitlements" .Entitlements}}{{else}}{{template "entitlements" .}}{{end}}</li>
    <li>{{if .Verification}}{{template "steps" .Verification}}{{else}}{{template "steps" .}}{{end}}</li>
    <li>{{if .Bypass}}{{template "steps" .Bypass}}{{else}}{{template "steps" .}}{{end}}</li>
    <li><input type="submit" value="submit"></li>
</ul>
{{end}}