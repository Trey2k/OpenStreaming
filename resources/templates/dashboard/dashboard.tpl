{{define "body"}}
    <div class="page" id="main">
        Welcome {{.StringOne}}
    </div>
    
    <div class="page" id="chatBot">
         <div class="switch">
            <label>
            Off
            <input type="checkbox" id="toggleBot">
            <span class="lever"></span>
            On
            </label>
        </div>
        <a class="waves-effect waves-light btn-small"id="checkEvents">Check Events</a>
    </div>

    
{{end}}