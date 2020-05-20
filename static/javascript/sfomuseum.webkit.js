var sfomuseum = sfomuseum || {};

sfomuseum.webkit = (function(){

    var self = {

	'init': function(){

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
	},
	
	'setAccessToken': function(token){
	    sfomuseum.console.log("Set OAuth2 access token");
	    document.body.setAttribute("data-oauth2-access-token", token);
	},

	'loadOEmbedURL': function(url){
	    
	    sfomuseum.console.log("Received oEmbed request for " + url);

	    var q_el = document.getElementById("oembed-url");

	    if (! q_el){
		sfomuseum.console.log("Missing oembed-url element.");
		return;
	    }
	    
	    var f_el = document.getElementById("oembed-fetch");
	    
	    if (! f_el){
		sfomuseum.console.log("Missing oembed-fetch element.");
		return;
	    }

	    q_el.value = url;
	    f_el.click();
	}
	
    };

    return self;

})();
