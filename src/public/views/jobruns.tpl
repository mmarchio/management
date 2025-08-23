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
                <td><a href="/jobruns/edit/{{$jobrun.ID}}">{{$jobrun.ID}}</a></td>
                <td>{{$jobrun.JobID}}</td>
                <td><a href="/jobruns/context/{{$jobrun.ID}}">context</a></td>
                <td><a href="/jobruns/settings/{{$jobrun.ID}}">settings</a></td>
                <td>{{if $jobrun.DispositionModel.Name}}{{$jobrun.DispositionModel.Name}}{{end}}</td>
                <td>{{$jobrun.Tokens}}</td>
                <td>{{$jobrun.Model.UpdatedAt}}</td>
                <td><a href="/jobruns/delete/{{$jobrun.ID}}">delete</a></td>
                <td><a href="/jobruns/run/{{$jobrun.ID}}">run</a></td>
            </tr>
            {{end}}
        </table>
            {{end}}
        {{end}}
        {{if eq .DisplayType "edit"}}
        <form action="/jobruns/save{{if .ID}}/{{.ID}}{{end}}" method="post">
            <input type="hidden" name="workflow_id" id="workflow_id" value="{{.Workflow.Model.ID}}">
            <ul>
                {{$nodes := .Nodes}}
                {{range $k, $v := .Steps}}
                <li>
                    <span style="float:left">{{if $v}}{{$v}}{{end}}</span>
                    <select style="float:right" name="{{$k}}" id="{{$k}}">
                        <option value="">Select Node</option>
                    {{range $kk, $vv := $nodes}}
                        <option value="{{$kk}}">{{$vv}}</option>
                    {{end}}
                    </select>
                {{end}}
                <li>{{template "element.submit" .}}</li>
            </ul>
        </form>
        <ul>
        </ul>
        {{end}}
    </body>
</html>