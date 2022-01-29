function newModule(moduleInfo) {
    switch(moduleInfo.Type) {
        case 1:
            return new AlertBox(moduleInfo.Top, moduleInfo.Left, moduleInfo.Width, moduleInfo.Height);
    }
}

function newAlertBox() {
    return new AlertBox((1080/2)-250, (1920/2)-250, 500, 500);
}

class Module {
    constructor(top, left, width, height) {
        this.top = top;
        this.left = left;
        this.width = width;
        this.Height = height;
    }

    sendEvent(event) {

    }
}

var index = 0;

class AlertBox extends Module {
    constructor(top, left, width, height) {
        super(top, left, width, height);
        this.id = index.toString();
        index++;
        $(".modules").append('<div class="alertBox module" id='+this.id+'></div>');
        var alertBox = $(".alertBox#"+this.id);
        alertBox.css("top", top);
        alertBox.css("left", left);
        alertBox.css("width", width);
        alertBox.css("height", height);
        if (editorMode) {
            alertBox.append(editorElements(this.id, "Alert Box"));
            alertBox.resizable({
                containment: ".overlay",
                scroll: false,
                handles: {
                    'nw': '.nw'+this.id,
                    'ne': '.ne'+this.id,
                    'sw': '.sw'+this.id,
                    'se': '.se'+this.id,
                    'n': '.n'+this.id,
                    'e': '.e'+this.id,
                    's': '.s'+this.id,
                    'w': '.w'+this.id
                }
            });
            
            alertBox.draggable({containment: ".overlay"});
        }
    }

    sendEvent(event) {
        switch (event.Type) {
            case 3:
                var alertBox = $(".alertBox#"+this.id);
                alertBox.append("<img src='"+event.Data.ProfilePicture+"' class=`viewerPFP`>");
                alertBox.append("<span class=`viewerName`>"+event.Data.DisplayName+"</span>");
        }
    }
}

function editorElements(id, title) {
    return `
    <span class="moduleTitle">`+title+`</span>
    <div class="ui-resizable-handle ui-resizable-nw nw`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-ne ne`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-sw sw`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-se se`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-n n`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-s s`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-e e`+id+`"></div>
    <div class="ui-resizable-handle ui-resizable-w w`+id+`"></div>
    `;
}