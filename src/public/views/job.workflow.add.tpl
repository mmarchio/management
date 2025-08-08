<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        <form action="/job/save/{{.Job.ID}}" method="POST">
            <ul>
                <li>
                    <select name="workflow_id" id="workflow_id">
                        <option value="">Select Workflow</option>
                    {{range .Workflows}}
                        <option value="{{.ID}}">{{.Name}}</option>
                    {{end}}
                    </select>
                </li>
                <li>{{template "element.submit" .}}</li>
            </ul>
        </form>
    </body>
</html>