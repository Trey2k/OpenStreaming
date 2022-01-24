class APIClient {
    constructor() {
        this.events = null;
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


}