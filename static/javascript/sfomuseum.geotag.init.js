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

    var feedback = document.createElement("ul");
    document.body.appendChild(feedback);

    var log = function(msg){

	console.log(msg);
	
	var item = document.createElement("li");
	item.appendChild(document.createTextNode(JSON.stringify(msg)));
	feedback.prepend(item);
    };
    
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
	    
	    var on_success = function(data){
		
		log("WRITE OKAY");
		log(data);
		
		try {
                    JSON.parse(data);
                }

		catch (e){
                    log("PARSE ERROR", e);
		    return;
		}

		log("OKAY");
		
		var wk_webview = document.body.getAttribute("data-enable-wk-webview");

		if (wk_webview == "true"){

		    log("WEBKIT IT UP...");

		    if (! sfomuseum.webkit.isAuth()){
			log("Not authenticated");
			return;
		    }
		    
		    try {
			webkit.messageHandlers.publishData.postMessage(data);
		    } catch(e) {
			log("SAD", e);
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
