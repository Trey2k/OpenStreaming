$(function()
{
  var output = document.getElementById("output");
  var input = document.getElementById("input");
  var ws;

  var print = function(message) {
    var d = document.createElement("div");
    d.innerHTML = message;
    output.appendChild(d);
  };

  $("#open").click(function() {
    var webSocket = new WebSocket("wss://weaselfoss.dev/ws");
    webSocket.onopen = function (event) {
      webSocket.send("Hello World!");
    }
    

  });

  $("#send").click(function() {
    if (!ws) {
      return false;
    }
    print("SEND: " + input.value);
    ws.send(input.value);
    return false;
  });

  $("#close").click(function() {
    if (!ws) {
      return false;
    }
    ws.close();
    return false;
  });
});