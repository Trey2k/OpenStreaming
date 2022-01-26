{{define "body"}}

<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script type="text/javascript" src="app.js"></script>

</head>
<body>
<table>
  <tr>
    <td valign="top" width="50%">
      <p>Click "Open" to create a connection to the server, 
      "Send" to send a message to the server and "Close" to close the connection. 
      You can change the message and send multiple times.
      </p>
        <button id="open">Open</button>
        <button id="close">Close</button>
        <p><input id="input" type="text" value="Hello world!">
        <button id="send">Send</button>
    </td>
    <td valign="top" width="50%">
      <div id="output"></div>
    </td>
  </tr>
</table>
</body>
</html>
{{end}}