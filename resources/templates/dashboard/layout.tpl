{{define "base"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/fontawesome.min.js"></script>
        <script type="text/javascript" src="/static/js/apiClient.js"></script>
        <script type="text/javascript" src="/static/js/dashboard/dashboard.js"></script>

        <link href="/static/css/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/materialize.min.css" rel="stylesheet">
        <link href="/static/css/dashboard/dashboard.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <main>
      <body id="base">

        <nav>
          <div class="nav-wrapper">
          <a data-target="slide-out" class="brand-logo center">Open<span class="tealText">Streaming</span></i></a>
          <ul id="nav-mobile" class="left">
          {{ if .LoggedIn }}
            <li><a data-target="slide-out" class="sidenav-trigger displayBlock button"><i class="fas fa-bars"></i></a></li>
          {{ end }}
          </ul>
        </div>
        </nav>

        <ul id="slide-out" class="sidenav">
        <li><a class="brand-logo">Open<span class="tealText">Streaming</span></a></li>
        {{ if .LoggedIn }}
          <li><div class="user-view">
            <div class="background">
              <img class="backgroundIMG" src="{{ .BackgroundPicture }}">
            </div>
            <a href="#user"><img class="circle" src="{{ .ProfilePicture }}"></a>
            <a href="#name"><span class="white-text name">{{ .DisplayName }} </span></a>
          </div></li>
          <li><a class="button" id="mainButton"><i class="fas fa-home"></i> Home</a></li>
          <li><a class="button" id="chatBotButton"><i class="fas fa-robot"></i> ChatBot</a></li>
          <li><a class="button" id="overlayButton" href="{{ .OverlayURL }}" target="_blank">Overlay</a></li>
        {{ end }}
        </ul>
        
        <div class="content">
            {{ template "body" . }}
        </div>
          
      </body>
    </main>
    <footer class="page-footer">
      <div class="container">
        <div class="row">
          <div class="col l6 s12">
            <h5 class="white-text">Open<span class="tealText">Streaming</span></h5>
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