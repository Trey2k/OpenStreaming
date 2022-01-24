(function(window) {
    var apiCLient = new APIClient;

    var refresh = function() {
        apiCLient.getEvents();
        requestAnimationFrame(animateImage);
    }

    imgLeft++;
    imgTop += slope;


    window.onload = function() {
        imgObj = document.getElementById('image');

        screenWidth = window.innerWidth;
        screenHeight = window.innerHeight;

        imgHeight = imgObj.offsetHeight;
        imgWidth = imgObj.offsetWidth;

        slope = (screenHeight - imgHeight) / (screenWidth - imgWidth);

        finalTop = screenHeight - imgHeight;
        finalLeft = screenWidth - imgWidth;

        requestAnimationFrame(refresh);
    };
})(window);