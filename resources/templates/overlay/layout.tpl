{{define "base"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/fontawesome.min.js"></script>
        <script type="text/javascript" src="/static/js/overlay/overlay.js"></script>

        <link href="/static/css/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/materialize.min.css" rel="stylesheet">
        <link href="/static/css/overlay/overlay.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <body id="min">
        <div class="content">
            {{ template "body" . }}
        </div>
    </body>
</html>

{{ end }}