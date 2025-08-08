<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        <form action="/api/jobruns" method="POST">
            <input type="hidden" name="id" id="id" value="{{if .ID}}{{.ID}}{{end}}">
            <input type="hidden" name="jobID" id="jobID" value="{{if .JobID}}{{.JobID}}{{end}}">
            <input type="hidden" name="settings" id="settings" value="{{if .Settings}}{{.Settings}}{{end}}">
            <input type="hidden" name="tokens" id="tokens" value="{{if .Tokens}}{{.Tokens}}{{end}}">
            <ul>
                <li><textarea name="context">{{if .Context}}{{.Context}}{{end}}</textarea></li>
                <li><input type="submit" value="submit></li>
            </ul>
        </form>
    </body>
</html>