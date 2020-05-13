var sfomuseum = sfomuseum || {};

sfomuseum.webkit = (function(){

    var _isauth = false;
    
    var self = {

	'setAuthOkay': function(){
	    _isauth = true;
	},

	'isAuth': function(){
	    return _isauth;
	},
	
    };

    return self;

})();
