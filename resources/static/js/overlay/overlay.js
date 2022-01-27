$(function()
{
  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get('id');
  const domain = window.location.host;
  var webSocket = new WebSocket("wss://"+domain+"/api/overlay/websocket?token="+id);
    
  webSocket.onopen = function (event) {
    webSocket.send("Hello World!");
  }

  window.onbeforeunload = function() {
    webSocket.onclose = function () {}; // disable onclose handler first
    webSocket.close();
  };
 
  webSocket.onmessage = function (event) {
    var e = JSON.parse(event.data);
    switch(e.Type) {
      case 2:
        $("#event").css({top: (1080/2)-500, left: (1920/2)-500});
        $("#event").append("<img src=\"/static/imgs/noOOP.png\" style=\"width: 500px; height: 500px;\">");
        $("#event").append("<p style=\"font-size: 50px;\">"+e.Data.MessageContent+"</p>");
    }
  }   

});

