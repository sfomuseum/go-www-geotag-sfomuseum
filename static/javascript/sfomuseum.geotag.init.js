window.addEventListener("load", function load(event){

    sfomuseum.console.log("INIT");
    
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
	    
	    var on_success = function(data){
		
		try {
                    var feature_collection = JSON.parse(data);
                }

		catch (e){
                    sfomuseum.console.log("PARSE ERROR", e);
		    return;
		}

		var access_token = document.body.getAttribute("data-oauth2-access-token");

		if (! access_token){
		    sfomuseum.console.log("MISSING ACCESS TOKEN");
		    return;
		}
		
		sfomuseum.console.log("GEOTAG TOKEN", access_token.length);
		
		var auth = {
		    token: access_token,
		};

		// https://github-tools.github.io/github/docs/3.2.3/Repository.html

		var gh = new GitHub(auth);

		var repo = gh.getRepo("sfomuseum-data", "sfomuseum-data-collection");
		
		var features = feature_collection["features"];
		var count = features.length;

		for (var i=0; i < count; i++){

		    var f = features[i];
		    var props = f["properties"];

		    var id = props["wof:id"];
		    var uri_args = {};

		    console.log(props);
		    
		    if (props["src:alt_label"]){

			var alt_parts = props["src:alt_label"].split("-");

			uri_args = {
			    "alt": true,
			    "source": alt_parts[0],
			    "function": alt_parts[1]
			};
		    }

		    var uri = whosonfirst.uri.id2relpath(id, uri_args);
		    var path = "data/" + uri;
		    
		    var branch = "master";
		    var content = f;

		    var message = "...";
		    var opts = {};

		    var cb = function(error, result, request){
			console.log("GH CB", error, result);
		    };
		    
		    repo.writeFile(branch, path, content, message, opts, cb)
		}

		return;
		
		var wk_webview = document.body.getAttribute("data-enable-wk-webview");

		if (wk_webview == "true"){

		    sfomuseum.console.log("WEBKIT IT UP...");

		    if (! sfomuseum.webkit.isAuth()){
			sfomuseum.console.log("Not authenticated");
			return;
		    }
		    
		    try {
			webkit.messageHandlers.publishData.postMessage(data);
		    } catch(e) {
			sfomuseum.console.log("SAD", e);
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
