{{define "body"}}
    <div class="page center" id="main">
        Welcome {{ .DisplayName }}
    </div>
    
    <div class="page" id="chatBot">
        
          <div class="row">
            <div class="col s12">
            <ul class="tabs">
                <li class="tab col s3"><a class="active" href="#chtBotMain">Main</a></li>
                <li class="tab col s3"><a href="#settings">Settings</a></li>
                <li class="tab col s3 disabled"><a href="#commands">Commands</a></li>
                <li class="tab col s3 disabled"><a href="#modules">Modules</a></li>
            </ul>
            </div>
            <div id="chtBotMain" class="col s12">
                
                <div class="row">
                    <div class="col s6">
                        
                    </div>
                    
                </div>

                  <div class="divider"></div>
                    <div class="section">
                        <div class="switch">
                            <label>
                                Off
                                <input type="checkbox" id="toggleBot">
                                <span class="lever"></span>
                                On
                            </label>
                        </div>
                    </div>
                    <div class="divider"></div>
                    <div class="section">
                        <div class="col s12"><a class="waves-effect waves-light btn-small"id="checkEvents">Check Events</a></div>
                    </div>
                    <div class="divider"></div>
                    <div class="section">
                        
                    </div>


            </div>
            
            </div>
            <div id="settings" class="col s12">
            
            </div>
            <div id="commands" class="col s12">
            
            </div>
            <div id="modules" class="col s12">
            
            </div>
        </div>
    </div>

    
{{end}}