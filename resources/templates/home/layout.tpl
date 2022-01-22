{{define "base"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/external/jquery/jquery-2.0.3.min.js"></script>
        <script type="text/javascript" src="/static/js/external/bootstrap/bootstrap.min.js"></script>

        <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-3.1.1/bootstrap.min.css">
        <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-3.1.1/bootstrap-theme.min.css">
        <link rel="stylesheet" type="text/css" href="/static/css/font-awesome-4.0.3/font-awesome.min.css">
        <link rel="stylesheet" type="text/css" href="/static/css/app.css">

        <title>{{ .Title }}</title>
    </head>
    <body>
            {{template "body" .}}
        </div>
    </body>
</html>

{{end}}