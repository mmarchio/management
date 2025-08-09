{{define "form.settings"}}
<table>
{{range $k, $v := .Debug.settings_model.template_model}}
    {{if eq $k "AvailableDispositionsArrayModel"}}
        {{range $i, $vv := $v}}
            <tr><td>$i</td></tr>
        {{end}}
    {{end}}
{{end}}
</table>
    <ul>
        {{if .}}
        <li>{{if .Prompt.SettingsModel.TemplateModel}}{{template "form.template" .Prompt.SettingsModel.TemplateModel}}{{else}}{{template "form.template" .}}{{end}}</li>
        <li>{{if .Prompt.SettingsModel.GlobalBypassModel}}{{template "steps" .Prompt.SettingsModel.GlobalBypassModel}}{{end}}</li>
        <li>
            <select name="workflow">
                <option value="">Select Workflow</option>
            {{range .Workflows}}
                <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
            </select>
        </li>
        <li><span>recurring</span><label class="switch"><input type="checkbox" name="recurring" id="recurring"{{if .Prompt.SettingsModel.RecurringModel.Value}}checked{{end}}><span class="slider round"></span></label></li>
        <li><input type="number" name="interval" id="interval" placeholder="interval" value="{{if .Prompt.SettingsModel.Interval}}{{.Prompt.SettingsModel.Interval}}{{end}}"></li>
        {{else}}
        <li>{{template "form.template" .Prompt.SettingsModel.TemplateModel}}</li>
        <li>{{template "steps" .Prompt.SettingsModel.GlobalBypassModel}}</li>
        <li>{{template "element.toggle" .Prompt.SettingsModel.RecurringModel}}</li>
        <li><input type="number" name="interval" id="interval" placeholder="interval" value="{{if .Prompt.SettingsModel.Interval}}{{.Prompt.SettingsModel.Interval}}{{end}}"></li>
        {{end}}
    </ul>
{{end}}