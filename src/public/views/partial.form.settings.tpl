{{define "form.settings"}}
    <ul>
        {{if .}}
        <li>{{if .Prompt.Settings.Template}}{{template "form.template" .Prompt.Settings.Template}}{{else}}{{template "form.template" .}}{{end}}</li>
        <li>{{if .Prompt.Settings.GlobalBypass}}{{template "steps" .Prompt.Settings.GlobalBypass}}{{end}}</li>
        <li>
            <select name="workflow">
                <option value="">Select Workflow</option>
            {{range .Workflows}}
                <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
            </select>
        </li>
        <li><span>recurring</span><label class="switch"><input type="checkbox" name="recurring" id="recurring"{{if .Prompt.Settings.Recurring}}checked{{end}}><span class="slider round"></span></label></li>
        <li><input type="number" name="interval" id="interval" placeholder="interval" value="{{if .Prompt.Settings.Interval}}{{.Prompt.Settings.Interval}}{{end}}"></li>
        {{else}}
        <li>{{template "form.template" .Prompt.Settings.Template}}</li>
        <li>{{template "steps" .Prompt.Settings.GlobalBypass}}</li>
        <li>{{template "element.toggle" .Prompt.Settings.Recurring}}</li>
        <li><input type="number" name="interval" id="interval" placeholder="interval" value="{{if .Prompt.Settings.Interval}}{{.Prompt.Settings.Interval}}{{end}}"></li>
        {{end}}
    </ul>
{{end}}