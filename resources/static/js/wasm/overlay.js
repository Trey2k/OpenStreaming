$(function()
{
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("/overlay.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
        const domain = window.location.host;
        if (startOverlay) {
            let err = startOverlay(domain, Token, EditorMode);
            if (err) {
                console.log("Error starting overlay:", err);
            }
        }
    });
});