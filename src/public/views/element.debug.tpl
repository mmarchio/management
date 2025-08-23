{{define "element.debug"}}
<table>
    <tr>
        <th>key</th>
        <th>value</th>
    </tr>
    {{range $k, $v := .}}
        <tr>
            <td>{{$k}}</td>
            <td>{{$v}}</td>
        </tr>
    {{end}}
</table>
{{end}}