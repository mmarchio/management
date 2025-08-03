<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        <form action="/prompts/settings/save/{{.ID}}" method="POST">
            <ul>
                <li>{{template "template" .Template}}</li>
                <li>{{template "steps" .GlobalBypass}}</li>
                <li>{{template "element.toggle" .Recurring}}</li>
                <li><input type="number" id="interval" name="interval" value="{{if .Interval}}{{.Interval}}{{end}}"></li>
                <li><input type="submit" value="submit">
            </ul>
        </form>
    </body>
</html>