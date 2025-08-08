{{define "submenu"}}
        <ul>
            <li><a href="/{{.Href}}">{{.Title}}</a></li>
            <li><a href="/{{.Href}}/new">New {{.Title}}</a></li>
            <li><a href="/{{.Href}}/list">List {{.Title}}</a></li>
        </ul>
{{end}}