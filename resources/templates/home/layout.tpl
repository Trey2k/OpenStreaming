{{define "base"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/external/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/external/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/external/fontawesome.min.js"></script>

        <link href="/static/css/external/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/external/materialize.min.css" rel="stylesheet">
        <link href="/static/css/openStreaming.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <body>
        <nav>
            <div class="nav-wrapper">
            <a href="#" class="brand-logo">Open<span class="tealText">Streaming</span></a>
            <ul id="nav-mobile" class="right hide-on-med-and-down">
            {{ if .LoggedIn }}
                <li><a class="button" href="sass.html">Link 1</a></li>
                <li><a class="button" href="badges.html">Link 2</a></li>
                <li><a class="button" href="collapsible.html">Link 3</a></li>
            {{ end }}
            </ul>
            </div>
        </nav>
        <div class="content">
            {{ template "body" . }}
        </div>
        <footer class="page-footer">
          <div class="container">
            <div class="row">
              <div class="col l6 s12">
                <h5 class="white-text">OpenStreaming</h5>
                <p class="grey-text text-lighten-4">Text about open streaming.</p>
              </div>
              <div class="col l4 offset-l2 s12">
              {{ if .LoggedIn }}
                <h5 class="white-text">Links</h5>
                <ul>
                  <li><a class="grey-text text-lighten-3 button" href="#!">Link 1</a></li>
                  <li><a class="grey-text text-lighten-3 button" href="#!">Link 2</a></li>
                  <li><a class="grey-text text-lighten-3 button" href="#!">Link 3</a></li>
                  <li><a class="grey-text text-lighten-3 button" href="#!">Link 4</a></li>
                </ul>
                {{ end }}
              </div>
            </div>
          </div>
        </footer>
    </body>
    
</html>

{{end}}