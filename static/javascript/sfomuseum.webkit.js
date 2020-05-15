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
	    
	    sfomuseum.console.log("GOT TOKEN");
	    document.body.setAttribute("data-oauth2-access-token", token);
	},
		
    };

    return self;

})();
