{{define "entitlements"}}
<details>
<summary>entitlements</summary>
<ul>
    <li>{{template "element.toggle" .YouTube}}</li>
    <li>{{template "element.toggle" .TikTok}}</li>
    <li>{{template "element.toggle" .Rumble}}</li>
    <li>{{template "element.toggle" .Patreon}}</li>
    <li>{{template "element.toggle" .Facebook}}</li>
</ul>
</details>
{{end}}