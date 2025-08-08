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
                <th>prompt</th>
                <th>workflow</th>
                <th>created at</th>
                <th>runs</th>
                <th>delete</th>
            </tr>
            {{range $job := .List}}
            <tr>
                <td>{{if $job.Model.ID}}{{$job.Model.ID}}{{end}}</td>
                <td>{{if $job.PromptID}}{{$job.PromptID}}{{end}}</td>
                <td>{{if $job.WorkflowID}}{{$job.WorkflowID}}{{else}}<a href="/job/workflow/add/{{$job.ID}}">add workflow</a>{{end}}</td>
                <td>{{if $job.CreatedAt}}{{$job.CreatedAt}}{{end}}</td>
                <td><a href="/jobruns/list/{{if $job.ID}}{{$job.ID}}{{end}}">runs</a></td>
                <td><a href="/jobruns/delete/{{$job.ID}}">delete</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
    </body>
</html>