{{define "overlay"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
        <script type="text/javascript" src="/static/js/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/fontawesome.min.js"></script>
        <script type="text/javascript">
          const Token = '{{.Token}}';
          const EditorMode = false;
        </script>
        <script type="text/javascript" src="/static/js/wasm/wasm_exec.js"></script>
        <script type="text/javascript" src="/static/js/wasm/overlay.js"></script>
        

        <link href="/static/css/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/materialize.min.css" rel="stylesheet">
        <link href="/static/css/jquery-ui.min.css" rel="stylesheet">
        <link href="/static/css/overlay/overlay.css" rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <body id="min">
        <div id="loadingScreen"></div>
        <div class="preloader-wrapper active" id="loadingIndicator">
          <div class="spinner-layer spinner-teal-only">
            <div class="circle-clipper left">
              <div class="circle"></div>
            </div><div class="gap-patch">
              <div class="circle"></div>
            </div><div class="circle-clipper right">
              <div class="circle"></div>
            </div>
          </div>
        </div>

        <div class="content">
            {{ template "overlayContent" . }}
        </div>
    </body>
</html>

{{ end }}

{{define "editor"}}

<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script type="text/javascript" src="/static/js/jquery-3.6.0.min.js"></script>
        <script type="text/javascript" src="/static/js/jquery-ui.min.js"></script>
        <script type="text/javascript" src="/static/js/materialize.min.js"></script>
        <script type="text/javascript" src="/static/js/fontawesome.min.js"></script>
        <script type="text/javascript">
            const Token = '{{.Token}}';
            const EditorMode = true;
        </script>
        <script type="text/javascript" src="/static/js/wasm/wasm_exec.js"></script>
        <script type="text/javascript" src="/static/js/wasm/editor.js"></script>
        <script type="text/javascript" src="/static/js/wasm/overlay.js"></script>

        <link href="/static/css/fontawesome.min.css" rel="stylesheet">
        <link href="/static/css/materialize.min.css" rel="stylesheet">
        <link href="/static/css/jquery-ui.min.css" rel="stylesheet">
        <link href="/static/css/overlay/overlay.css" rel="stylesheet">
        <link href="/static/css/overlay/editor.css" rel="stylesheet">
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">

        <title>{{ .Title }}</title>
    </head>
    <body id="min">
        <ul id="slide-out" class="sidenav">
        <li><a class="brand-logo">Open<span class="tealText">Streaming</span></a></li>
          <li><div class="user-view">
            <div class="background">
              <img class="backgroundIMG" src="{{ .BackgroundPicture }}">
            </div>
            <a href="#user"><img class="circle" src="{{ .ProfilePicture }}"></a>
            <a href="#name"><span class="white-text name">{{ .DisplayName }} </span></a>
          </div></li>
          <li><a class="button" id="mainButton" href="/dashboard"><i class="fas fa-home"></i> Home</a></li>
          <li><a class="button" id="chatBotButton" href="/dashboard#chatBot"><i class="fas fa-robot"></i> ChatBot</a></li>
          <li><a class="button active" id="overlayButton" href="#"><i class="fas fa-magic"></i> Overlay Editor</a></li>
        </ul>

        <div id="loadingScreen"></div>
        <div class="preloader-wrapper active" id="loadingIndicator">
          <div class="spinner-layer spinner-teal-only">
            <div class="circle-clipper left">
              <div class="circle"></div>
            </div><div class="gap-patch">
              <div class="circle"></div>
            </div><div class="circle-clipper right">
              <div class="circle"></div>
            </div>
          </div>
        </div>

        <div class="content">
            {{ template "editorContent" . }}
        </div>
    </body>
</html>

{{ end }}