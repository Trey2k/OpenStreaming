{{define "overlayContent"}}
  <div class="overlay">
    <div class="modules">
    </div>
  </div>
{{end}}

{{define "editorContent"}}
  <div class="overlay">
      <div class="control">
        <a data-target="slide-out" class="sidenav-trigger btn-floating btn-large waves-effect waves-light red" id="menuButton"><i class="material-icons">menu</i></a>
        <a class="btn-floating btn-large waves-effect waves-light blue" id="fullScreen"><i class="material-icons">aspect_ratio</i></a>
        <a class="btn-floating btn-large waves-effect waves-light green" id="save"><i class="material-icons">save</i></a>

        <div class="fixed-action-btn">
                <a class="btn-floating btn-large waves-effect waves-light red" id="newModule"><i class="material-icons">add</i></a>
            <ul>
                <li><a class="btn-floating yellow" id="addAlertBox"><i class="material-icons">add_alert</i></a></li>
            </ul>
        </div>
    </div>
    <div class="modules">
    </div>
  </div>
{{end}}