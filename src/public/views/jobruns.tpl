<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        {{if .DisplayType}}
            {{if eq .DisplayType "none"}}
            {{end}}
            {{if eq .DisplayType "new"}}
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>id</th>
                <th>job id</th>
                <th>context</th>
                <th>settings</th>
                <th>disposition</th>
                <th>tokens</th>
                <th>last updated</th>
                <th>delete</th>
                <th>run</th>
            </tr>
            {{range $jobrun := .List}}
            <tr>
                <td>{{$jobrun.ID}}</td>
                <td>{{$jobrun.JobID}}</td>
                <td><a href="/jobruns/context/{{$jobrun.ID}}">context</a></td>
                <td><a href="/jobruns/settings/{{$jobrun.ID}}">settings</a></td>
                <td>{{if $jobrun.Disposition.Name}}{{$jobrun.Disposition.Name}}{{end}}</td>
                <td>{{$jobrun.Tokens}}</td>
                <td>{{$jobrun.Model.UpdatedAt}}</td>
                <td><a href="/jobruns/delete/{{$jobrun.ID}}">delete</a></td>
                <td><a href="/jobruns/run/{{$jobrun.ID}}">run</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>