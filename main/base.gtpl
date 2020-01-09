<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Items}}
            <div>raka
                {{ . }}
            </div>
        {{else}}
            <div>
                <strong>Tidak ada menu</strong>
            </div>
        {{end}}
    </body>
</html>