var sfomuseum = sfomuseum || {};

sfomuseum.console = (function(){

    var _feedback;

    var self = {
	
	'log': function(msg){

	    if (! _feedback){
		_feedback = document.createElement("ul");
		document.body.appendChild(_feedback);
	    }

	    if (typeof(msg) != "string"){
		msg = JSON.stringify(msg);
	    }
	    
	    var item = document.createElement("li");
	    item.appendChild(document.createTextNode(msg));	    
	    _feedback.prepend(item);

	    console.log(msg);
	},
	
    };

    return self;

})();    
