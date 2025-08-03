<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        {{if .DisplayType}}
            {{if eq .DisplayType "none"}}
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>id</th>
                <th>name</th>
                <th>min duration</th>
                <th>max duration</th>
                <th>advertisement duration</th>
                <th>entitlements</th>
                <th>verification</th>
                <th>bypass</th>
                <th>delete</th>
            </tr>
            {{range $disposition := .List}}
            <tr>
                <td><a href="/disposition/{{$disposition.ID}}">{{$disposition.ID}}</a></td>
                <td>{{$disposition.Name}}</td>
                <td>{{$disposition.MinDuration}}</td>
                <td>{{$disposition.MaxDuration}}</td>
                <td>{{$disposition.AdvertisementDuration}}</td>
                <td><a href="/dispositions/entitlements/{{$disposition.ID}}">entitlements</a></td>
                <td><a href="/dispositions/verification/{{$disposition.ID}}">verification</a></td>
                <td><a href="/dispositions/bypass/{{$disposition.ID}}">bypass</a></td>
                <td><a href="/dispositions/delete/{{$disposition.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
            {{if eq .DisplayType "new"}}
        <form action="/dispositions/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.disposition" .}}
        </form>
            {{end}}
        {{end}}
    </body>
</html>