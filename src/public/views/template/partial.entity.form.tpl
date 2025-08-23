{{define "entity.form"}}
        {{if eq .ContentType "comfy"}}
        {{template "form.entity" .}}
        {{end}}
        {{if eq .ContentType "prompt"}}
        {{template "form.prompt" .}}
        {{end}}
        {{if eq .ContentType "disposition"}}
        {{template "form.disposition .}}
        {{end}}
        {{if eq .ContentType "systemprompt"}}
        {{template "form.systemprompt" .}}
        {{end}}
        {{if eq .ContentType "job"}}
        {{template "form.job" .}}
        {{end}}
        {{if eq .ContentType "jobrun"}}
        {{template "form.jobrun" .}}
        {{end}}
        {{if eq .ContentType "entitlements"}}
        {{template "form.entitlements" .}}
        {{end}}
        {{if eq .ContentType "bypass"}}
        {{template "form.steps" .}}
        {{end}}
        {{if eq .ContentType "verification"}}
        {{template "form.steps" .}}
        {{end}}
        {{if eq .ContentType "global_bypass"}}
        {{template "form.steps" .}}
        {{end}}
        {{if eq .ContentType "template"}}
        {{template "form.template" .}}
        {{end}}
        {{if eq .ContentType "settings"}}
        {{template "form.settings" .}}
        {{end}}
{{end}}