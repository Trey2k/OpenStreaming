function showChatBot() {
    $("#main").hide();
    $("#mainButton").removeClass("active");
    $("#chatBot").show();
    $("#chatBotButton").addClass("active");
}

function showMain() {
    $("#chatBot").hide();
    $("#chatBotButton").removeClass("active");
    $("#main").show();
    $("#mainButton").addClass("active");
}

$(function()
{

    var apiClient = new APIClient;
    var botOn = false;
    showMain();

    $(".button").click(function() {
        switch(this.id) {
            case "chatBotButton":
                showChatBot();
                console.log("ChatBot");
                break;
            case "mainButton":
                showMain();
                console.log("main");
                break;
        }
    });

    $("#toggleBot").click(function() {
        apiClient.toggleBot();
        botOn = !botOn;
        if (botOn) {
            apiClient.onEvent(function(event) {
                switch(event.Type) {
                case TwitchMessageEvent: {
                    var message = event.Data.MessageContent;
                    console.log(event.Data.UserDisplay+":", message);
                }
                }
            });
        }
    });

    

    $("#checkEvents").click(function() {
        events = apiClient.getEvents();
        if (events != null) {
            for (var i = 0; i < events.length; i++) {
                console.log(events[i]);
            }
        }
        

    });
});

