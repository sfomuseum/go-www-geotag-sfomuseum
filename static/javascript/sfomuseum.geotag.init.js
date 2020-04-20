window.addEventListener("load", function load(event){

    var map = geotag.maps.getMapById("map");

    if (! map){
	console.log("Failed to retrieve map");
	return;
    }
    
    var layers_control = new L.Control.Layers({
	catalog: sfomuseum.maps.catalog,
    });

    map.addControl(layers_control);

    // https://stackoverflow.com/questions/50229935/wkwebview-get-javascript-errors

    window.onerror = (msg, url, line, column, error) => {
	const message = {
	    message: msg,
	    url: url,
	    line: line,
	    column: column,
	    error: JSON.stringify(error)
	}
	
	if (window.webkit) {
	    window.webkit.messageHandlers.error.postMessage(message);
	} else {
	    console.log("Error:", message);
	}
    };
});
