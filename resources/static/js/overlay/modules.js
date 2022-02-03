function newModule(moduleInfo) {
    switch(moduleInfo.Type) {
        case 1:
            return new AlertBox(moduleInfo.Top, moduleInfo.Left, moduleInfo.Width, moduleInfo.Height, false);
    }
}

function newAlertBox() {
    return new AlertBox((1080/2)-250, (1920/2)-250, 500, 500, true);
}

class Module {
    Top = 0;
    Left = 0;
    Width = 0;
    Height = 0
    Events = [];
    Busy = false;
    EventHandler= null;

    constructor(top, left, width, height, isNew) {
        this.Top = top;
        this.Left = left;
        this.Width = width;
        this.Height = height;
        this.Events = [];
        this.Busy = false;
        this.IsNew = isNew;
        this.EventHandler = setInterval(this.eventHandler.bind(this), 1);
    }

    getTop() {
        return this.Top;
    }

    getLeft() {
        return this.Left;
    }

    getWidth() {
        return this.Width;
    }

    getHeight() {
        return this.Height;
    }

    getNew() {
        return this.IsNew;
    }

    update() {

    }

    sendEvent(event) {
        console.log("Event: " + event.Type);
        this.Events.push(event);
    }

    eventHandler() {
        if (this.Events.length > 0 &&  this.Busy == false) {
            let event = this.Events.pop();
            switch (event.Type) {
                case 3:
                    this.followEvent(event);
            }
        }
    }

    destroy() {
        clearInterval(this.EventHandler);
    }

    getType() {
        return 0;
    }

    followEvent(event) {
    }

}

var index = 0;

class AlertBox extends Module {
    id = "";
    module = null;
    alertBox = null;
    constructor(top, left, width, height, isNew) {
        super(top, left, width, height, isNew);
        this.id = index.toString();
        index++;
        $(".modules").append('<div class="module" id='+this.id+'><div class="alertBox" id='+this.id+'></div></div>')
        this.module = $(".module#"+this.id);
        this.alertBox = this.module.children(".alertBox");
        this.module.css("top", top).css("left", left).css("width", width).css("height", height);
        
        if (editorMode) {
            this.module.prepend(editorElements(this.id, "Alert Box"));
            this.module.resizable({
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
            
            this.module.draggable({containment: ".overlay"});
        }
    }

    update() {
        this.Top = this.module.position().top;
        this.Left = this.module.position().left;
        this.Width = this.module.width();
        this.Height = this.module.height();
    }


    destroy() {
        super.destroy();
        this.module.remove();
    }

    getType() {
        return 1;
    }

    followEvent(event) {
        this.Busy = true;
            
        this.alertBox.append("<img src='"+event.Data.ProfilePicture+"' class='viewerPFP alertInfo'>");
        this.alertBox.append("<span class='viewerName alertInfo'>Thanks for following "+event.Data.DisplayName+"!</span>");
        setTimeout(this.clear.bind(this), 15 * 1000);
    }

    clear() {
        
        this.alertBox.empty();
        this.Busy = false;
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