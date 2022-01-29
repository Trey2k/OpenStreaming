let fullScreen = false;

function initEditor(modules){
  $('.sidenav').sidenav();
  
  const elems = document.querySelectorAll('.fixed-action-btn');
  const instances = M.FloatingActionButton.init(elems, {
    direction: 'left',
    hoverEnabled: false
  });

  $("#addAlertBox").click(function(){
      modules.push(newAlertBox());
  });

  $("#fullScreen").click(function(){
    
    if (!fullScreen) {
      openFullscreen();
    } else {
      closeFullscreen();
    }
    fullScreen = !fullScreen;
  });
}


function openFullscreen() {
  const elem = document.documentElement;
  if (elem.requestFullscreen) {
    elem.requestFullscreen();
  } else if (elem.webkitRequestFullscreen) { 
    elem.webkitRequestFullscreen();
  } else if (elem.msRequestFullscreen) { 
    elem.msRequestFullscreen();
  }
}

function closeFullscreen() {
  if (document.exitFullscreen) {
    document.exitFullscreen();
  } else if (document.webkitExitFullscreen) { /* Safari */
    document.webkitExitFullscreen();
  } else if (document.msExitFullscreen) { /* IE11 */
    document.msExitFullscreen();
  }
}