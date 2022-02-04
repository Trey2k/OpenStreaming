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
        
        

        <div class="fixed-action-btn" id="actions" style="display: none;">
                <a class="btn-floating btn-large waves-effect waves-light red" id="quickAction"><i class="material-icons">flash_on</i></a>
            <ul>
            <li><a class="btn-floating btn-large waves-effect waves-light green" id="save"><i class="material-icons">save</i></a></li>
            <li><a class="btn-floating btn-large waves-effect waves-light blue" id="fullScreen"><i class="material-icons">aspect_ratio</i></a></li>
            </ul>
        </div>

          <div class='context-trigger' data-target='contextMenu'></div>
          
          <ul id='contextMenu' class='dropdown-content'>
            <li class="divider" tabindex="-1"></li>
            <li><a class='addMenu-trigger' data-target='addMenu'><i class="material-icons">add</i>Add</a></li>
          </ul>
          
          <ul id='addMenu' class='dropdown-content addMenuConent'>
            <li><a id="addAlertBox"><i class="material-icons">add_alert</i>Alert Box</a></li>
          </ul>

    </div>
    <div class="modules">
    
  </div>
{{end}}