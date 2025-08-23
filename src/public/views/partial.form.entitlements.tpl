{{define "entitlements"}}
<details>
<summary>entitlements</summary>
<ul>
    <li>{{template "element.toggle" .YouTubeModel}}</li>
    <li>{{template "element.toggle" .TikTokModel}}</li>
    <li>{{template "element.toggle" .RumbleModel}}</li>
    <li>{{template "element.toggle" .PatreonModel}}</li>
    <li>{{template "element.toggle" .FacebookModel}}</li>
</ul>
</details>
{{end}}