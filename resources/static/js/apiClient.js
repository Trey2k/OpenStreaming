class APIClient {
    constructor() {
        this.events = null;
        this.on = false;
    }

    getEvents() {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", "/api/getEvents", false); // false for synchronous request
        xmlHttp.send(null);
        this.events = JSON.parse(xmlHttp.responseText);
        return this.events
    }

    toggleBot() {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", "/api/toggleBot", false); // false for synchronous request
        xmlHttp.send(null);
    }

    onEvent(eventHandler) {
        var client = this;
        setInterval(function() {
            var events = client.getEvents();
            if (events != null) {
                for (var i = 0; i < events.length; i++) {
                    eventHandler(events[i]);
                }
            }
        }, 1000);   
    }
}

const InvalidEvent = 0, TestEvent = 1, TwitchMessageEvent = 2;
