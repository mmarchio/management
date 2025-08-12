{{define "form.step"}}
<ul>
    <li><input type="text" name="name" id="name" placeholder="name" value="{{if .Name}}{{.Name}}{{end}}"></li>
    <li><input type="number" name="order" id="order" placeholder="order" value="{{if .Order}}{{.Order}}{{end}}"></li>
    <li>
        <select name="disposition_id" id="disposition_id">
            <option value="">Select Disposition</option>
        {{range .Dispositions}}
            <option value="{{.ID}}">{{.Name}}</option>
        {{end}}
        </select>
    </li>
    <li>{{template "element.toggle" .Enabled}}</li>
    <li>{{template "element.toggle" .Bypass}}</li>
    <li>{{template "element.submit" .}}</li>
</ul>
{{end}}

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
        <form action="/step/save{{if .ID}}/{{.ID}}{{end}}" method="POST">
        {{template "form.step" .}}
        </form>
            {{end}}
            {{if eq .DisplayType "list"}}
        <table>
            <tr>
                <th>id</th>
                <th>name</th>
                <th>order</th>
                <th>dispositionID</th>
                <th>enabled</th>
                <th>bypass</th>
            </tr>
            {{range .List}}
            <tr>
                <td><a href="/step/edit{{if .ID}}/{{.ID}}{{end}}">{{if .ID}}{{.ID}}{{end}}</a></td>
                <td>{{.Name}}</td>
                <td>{{.Order}}</td>
                <td>{{.DispositionID}}</td>
                <td>{{if .Enabled.Value}}true{{else}}false{{end}}</td>
                <td>{{if .Bypass.Value}}true{{else}}false{{end}}</td>
                <td><a href="/step/delete{{if .ID}}/{{.ID}}{{end}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>