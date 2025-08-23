<!DOCTYPE html>
<html>
    {{template "head" .}}
    <body>
        {{template "menu" .}}
        {{template "submenu" .Menu}}
        <pre style="float:left; clear:both;">{{.}}</pre>
    </body>
</html>