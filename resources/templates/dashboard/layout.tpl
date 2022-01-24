{{define "base"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/external/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/external/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/external/fontawesome.min.js"></script>
        <script type="text/javascript" src="/static/js/apiClient.js"></script>
        {{ if .CustomJS }}
        <script type="text/javascript" src="{{ .CustomJSPath }}"></script>
        {{ end }}

        <link href="/static/css/external/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/external/materialize.min.css" rel="stylesheet">
        <link href="/static/css/openStreaming.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <main>
      <body id="base">
          <nav>
              <div class="nav-wrapper">
              <a href="#" class="brand-logo">Open<span class="tealText">Streaming</span></a>
              <ul id="nav-mobile" class="right hide-on-med-and-down">
              {{ if .LoggedIn }}
                  <li><a class="button" id="mainButton"><i class="fas fa-home"></i></a></li>
                  <li><a class="button" id="chatBotButton"><i class="fas fa-robot"></i></a></li>
              {{ end }}
              </ul>
              </div>
          </nav>
          <div class="content">
              {{ template "body" . }}
          </div>
          
      </body>
    </main>
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
            {{ end }}
            <ul>
              {{ if .LoggedIn }}
              <li><a class="grey-text text-lighten-3 button" href="#!">Link 1</a></li>
              <li><a class="grey-text text-lighten-3 button" href="#!">Link 2</a></li>
              <li><a class="grey-text text-lighten-3 button" href="#!">Link 3</a></li>
              <li><a class="grey-text text-lighten-3 button" href="#!">Link 4</a></li>
              {{ end }}
            </ul>
          </div>
        </div>
      </div>
    </footer>
</html>

{{end}}

{{define "min"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/external/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/external/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/external/fontawesome.min.js"></script>
        {{ if .CustomJS }}
        <script type="text/javascript" src="{{ .CustomJSPath }}"></script>
        {{ end }}

        <link href="/static/css/external/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/external/materialize.min.css" rel="stylesheet">
        <link href="/static/css/openStreaming.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <body id="min">
        <div class="content">
            {{ template "body" . }}
        </div>
    </body>
</html>

{{end}}