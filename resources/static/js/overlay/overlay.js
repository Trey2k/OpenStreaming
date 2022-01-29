var token = "";
var editorMode = false;

function setToken(t){
  token = t
}

function initOverlay(overlay) {
  var modules = [];
  for(var i = 0; i < overlay.Modules.length; i++) {
    modules.push(newModule(overlay.Modules[i]));
  }

  return modules;
}

$(function()
{
  const domain = window.location.host;
  var webSocket = new WebSocket("wss://"+domain+"/api/overlay/websocket?token="+token);
  var ovarlay = null;
  var modules = [];
    
  webSocket.onopen = function (event) {
    console.log("WebSocket opened");
    webSocket.send(JSON.stringify({
      "Type": "getOverlay",
      "Overlay": null
    }));
  }

  window.onbeforeunload = function() {
    webSocket.onclose = function () {}; // disable onclose handler first
    webSocket.close();
  };
 
  webSocket.onmessage = function (event) {
    console.log(event.data);
    var e = JSON.parse(event.data);
    switch(e.Type) {
      case "Return":
        ovarlay = e.Overlay
        modules = initOverlay(ovarlay);
      default:
        for (let i = 0; i < modules.length; i++) {
          modules[i].sendEvent(e);
        }
    }
  }   

  if (editorMode) {
    initEditor(modules);
  }

});

