var sfomuseum = sfomuseum || {};

sfomuseum.webkit = (function(){

    var _isauth = false;
    
    var self = {

	'setAuthOkay': function(){
	    _isauth = true;
	},

	'setAccessToken': function(token){
	    
	    sfomuseum.console.log("GOT TOKEN");
	    document.body.setAttribute("data-oauth2-access-token", token);
	},
	
	'isAuth': function(){
	    return _isauth;
	},
	
    };

    return self;

})();
