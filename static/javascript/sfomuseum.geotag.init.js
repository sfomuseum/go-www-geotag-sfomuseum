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

    var save = document.getElementById("writer-save");

    if (save){

	save.onclick = function(e){

	    var camera = geotag.camera.getCamera();
	    
            if (! camera){
		console.log("Unable to retrieve camera");
		return false;
	    }
	    
            var uri = document.body.getAttribute("data-geotag-uri", uri);
	    
	    if (! uri){
		console.log("Missing data-geotag-uri attribute");
		return false;
	    }
	    
            var fov = camera.getFieldOfView();
	    
	    var on_success = function(rsp){
		
		console.log("WRITE OKAY", rsp);

		var wk_webview = document.body.getAttribute("data-enable-wk-webview");

		if (wk_webview == "true"){

		    console.log("WEBKIT IT UP...");
		    
		    try {
			var body = JSON.stringify(rsp);
			
			var wk_rsp = webkit.messageHandlers.publishData.postMessage(body);
			console.log("WEBKIT RESPONSE", wk_rsp);
		    } catch(e) {
			console.log("SAD", e);
		    }

		    console.log("DONE");
		}
		
	    };
	    
            var on_error = function(err){
		console.log("WRITE ERROR", err);
	    };
	    
            geotag.writer.write_geotag(uri, fov, on_success, on_error);
	    return false;
	};
    }
    
});
