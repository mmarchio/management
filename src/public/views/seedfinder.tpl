{{define "form.seedfinder"}}
<ul>
    <li>
        {{$workflow_id := ""}}
        {{if .dt.WorkflowID}}{{$workflow_id = .dt.WorkflowID}}{{end}}
        <select name="workflow_id" id="workflow_id">
            <option value="">Select Workflow</option>
        {{range .dt.Workflows}}
            {{if eq $workflow_id .ID}}
            <option value="{{.ID}}" selected="selected">{{.Name}}</option>
            {{else}}
            <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
        {{end}}
        </select>
    </li>
    <li>
        {{$step_id := ""}}
        {{if .dt.StepID}}{{$step_id = .dt.StepID}}{{end}}
        <select name="step_id" id="step_id">
            <option value="">Select Step</option>
        {{range .dt.Steps}}
            <option value="{{.ID}}" {{if eq $step_id .ID}} selected="selected"{{end}}>{{.Name}}</option>
        {{end}}
        </select>
    </li>
    <li>
        {{$job_id := ""}}
        {{if .dt.JobID}}{{$job_id = .dt.JobID}}{{end}}
        <select name="job_id" id="job_id">
            <option value="">Select Job</option>
        {{range .dt.Jobs}}
            <option value="{{.ID}}" {{if eq $job_id .ID}} selected="selected"{{end}}>{{.Name}}</option>
        {{end}}
        </select>
    </li>
    <li><input type="number" id="number" name="number" value="limit"></li>
    <li>{{template "element.submit" .}}</li>
</ul>
{{end}}

<!DOCTYPE html>
<html>
    <!--Template head /-->
    {{template "head" .}}
    <body>
        <!-- Template menu /-->
        {{template "menu" .}}
        <!-- Template submenu /-->
        <ul>
            <li><a href="/seedfinder/configure">configure</a></li>
        </ul>
{{if .dt.DisplayType}}
    {{if eq .dt.DisplayType "none"}}
    <span>none</span>
            <!-- displaytype none /-->
    {{end}}
    {{if eq .dt.DisplayType "configure"}}
    <span>configure</span>
            <!-- displaytype new /-->
        <form action="/seedfinder/run" method="POST" name="new">
        <!-- template form.seedfinder /-->
        {{template "form.seedfinder" .}}
        </form>
    {{end}}
    {{if eq .dt.DisplayType "run"}}
    {{$job_id := ""}}
    {{if .dt.JobID}}{{$job_id = .dt.JobID}}{{end}}
    {{$step_id := ""}}
    {{if .dt.StepID}}{{$step_id = .dt.StepID}}{{end}}
    <ul>
    {{range $k, $v := .}}
    {{if eq $k "dt"}}{{continue}}{{end}}
    <li>
        <video width="512" height="512" controls>
            <source src="http://172.17.0.1:9002/view?filename={{$v}}&type=output" type="video/mp4">
        </video>
        <a href="/seedfinder/set/{{$job_id}}/{{$step_id}}/{{$k}}">{{$k}}</a>
    </li>
    {{end}}
    {{end}}
    </ul>
{{end}}
    </body>
</html>