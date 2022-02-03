var token = "";
function setToken(t){
  token = t
}

function initOverlay(overlay, webSocket) {
  overlay.Mods = [];
  console.log("Modules", overlay.Modules);
  for(const i in overlay.Modules) {
    console.log("Module", overlay.Modules[i]);
    overlay.Mods.push(newModule(overlay.Modules[i]));
  }
}

$(function()
{
  const domain = window.location.host;
  var webSocket = new WebSocket("wss://"+domain+"/api/overlay/websocket?token="+token);
  var overlay = null;
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
    var e = JSON.parse(event.data);
    switch(e.Type) {
      case "return":
        overlay = e.Overlay
        initOverlay(overlay, webSocket);
        if (editorMode) {
          console.log("Editor mode enabled");
          initEditor(overlay, webSocket);
        }
        break;
      default:
        for (let i = 0; i < modules.length; i++) {
          modules[i].sendEvent(e);
        }
    }
  }   

  

});

