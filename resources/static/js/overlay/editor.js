let fullScreen = false;
let contextTarget = null;
let conextLocation = {};
const editorMode = true;

let modules = null;

function initEditor(overlay, websocket){
  modules = overlay.Mods;
  $('.sidenav').sidenav();

  let dropDown = $('.context-trigger').dropdown({
		constrain_width: true, // Does not change width of dropdown to that of the activator
		gutter: 0, // Spacing from edge
		belowOrigin: true, // Displays dropdown below the button
		alignment: 'left' // Displays dropdown with edge aligned to the left of button
  });
  
  let addMenu = $('.addMenu-trigger').dropdown({
    inDuration: 300,
		outDuration: 225,
		constrain_width: true, // Does not change width of dropdown to that of the activator
		hover: true, // Activate on hover
		gutter: 0, // Spacing from edge
		belowOrigin: true, // Displays dropdown below the button
		alignment: 'left' // Displays dropdown with edge aligned to the left of button
  });

  const elems = document.querySelectorAll('.fixed-action-btn');
  const instances = M.FloatingActionButton.init(elems, {
    direction: 'left',
    hoverEnabled: true
  });

  $("#addAlertBox").click(function(){
      modules.push(new AlertBox(conextLocation.y-175, conextLocation.x-242, 350, 485, true));
  });

  $("#fullScreen").click(function(){
    
    if (!fullScreen) {
      openFullscreen();
    } else {
      closeFullscreen();
    }
    fullScreen = !fullScreen;
  });

  $("#save").click(function(){
    console.log("Saving");
    let modulesToSend = [];
    for(let i = 0; i < modules.length; i++){
      modules[i].update();
      modulesToSend.push({Top: modules[i].getTop(), 
                          Left: modules[i].getLeft(), 
                          Width: modules[i].getWidth(), 
                          Height: modules[i].getHeight(), 
                          Type: modules[i].getType(), 
                          ID: parseInt(modules[i].id),
                          IsNew: modules[i].getNew()
      });
    }

    let event = {
      Type: "updateOverlay",
      Modules: modulesToSend
    }
    websocket.send(JSON.stringify(event));

  });

    document.addEventListener('contextmenu', function(e) {
      dropDown.dropdown('close');
      hideModuleContext();
      if($(e.target).is('.module')||$(e.target).is('.alertBox')) {
        contextTarget = $(e.target);
        showModeulContext();
      } else if (contextTarget != null) {
        contextTarget = null;
      }

      const x = e.clientX+5;
      const y = e.clientY+5;
      conextLocation = {x: x, y: y};

      dropDown.css('top',`${y}px`).css('left',`${x}px`);

      dropDown.dropdown('open');
      e.preventDefault();

    }, false);
  
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

function showModeulContext() {
    $("#contextMenu").prepend(`<li class="contextModule" id="delete" onclick="deleteContextModule()"><a class="red-text text-darken-1"><i class="material-icons">cancel</i>Delete</a></li>`);
    
}

function hideModuleContext() {
    $("#contextMenu").children().remove('.contextModule');
}

function deleteContextModule(){
 
  if (contextTarget != null) {
    for (let i = 0; i < modules.length; i++) {
      if (modules[i].id == contextTarget.attr('id')) {
        let module = modules[i];
        modules.splice(i, 1);
        module.destroy();
        break;
      }
    }
    contextTarget = null;
  }
}