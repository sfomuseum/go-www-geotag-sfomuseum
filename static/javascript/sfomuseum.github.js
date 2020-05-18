// this is not a fully-fledge github API client - it only does enough to
// add or update files per these methods:
// https://developer.github.com/v3/repos/contents/#get-contents
// https://developer.github.com/v3/repos/contents/#create-or-update-a-file
// (2020518/thisisaaronland)

// I tried to use this but I could never get it to return an actual SHA
// value for a file and it didn't look like I could assign it for update
// https://github-tools.github.io/github/docs/3.2.3/
// (20200518/thisisaaronland)

var sfomuseum = sfomuseum || {};

sfomuseum.github = (function(){

    var _access_token;

    var self = {

	'setAccessToken': function(token){
	    _access_token = token;
	},
	
	'updateFile': function(branch, path, args, on_success, on_error) {

	    var on_success_sha = function(sha){
		args["sha"] = sha;
		args["message"] = args["message"].replace("Add", "Update");
		self.writeFile(branch, path, args, on_success, on_error);
	    };

	    var on_error_sha = function(err){
		self.writeFile(branch, path, args, on_success, on_error);			
	    };		    
		    
	    self.getSha(branch, path, on_success_sha, on_error_sha);
	},
	
	'getSha': function(branch, path, on_success, on_error){
	    
	    var on_load = function(rsp){

		var target = rsp.target;

		if (target.readyState != 4){
		    return;
		}

		var status_code = target['status'];
		var status_text = target['statusText'];

		if ((status_code < 200) || (status_code > 299)){
		    on_error({'code': status_code, 'message': status_text});
		    return;
		}

		var raw = target['responseText'];

		try {
                    var data = JSON.parse(raw);
                }

		catch (e){
		    on_error(e);
		    return;
		}
		
		on_success(data.sha);
	    };
	    
	    var req = new XMLHttpRequest();
	    
	    req.addEventListener("load", on_load);
	    req.addEventListener("error", on_error);

	    var uri = "https://api.github.com/repos/sfomuseum-data/sfomuseum-data-collection/contents/" + path + "?ref=" + branch;
	    req.open("GET", uri, true);

	    req.setRequestHeader("Accept", "application/vnd.github.v3.json");
	    req.setRequestHeader("Authorization", "token " + _access_token);	    
	    req.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

	    req.send(null);    
	},

	'writeFile': function(branch, path, args, on_success, on_error){

	    var form_data = JSON.stringify(args);
	    
	    var on_load = function(rsp){

		var target = rsp.target;

		if (target.readyState != 4){
		    return;
		}

		var status_code = target['status'];
		var status_text = target['statusText'];

		if ((status_code < 200) || (status_code > 299)){
		    on_error({'code': status_code, 'message': status_text});
		    return;
		}

		var raw = target['responseText'];
		on_success(raw);
	    };
	    
	    var req = new XMLHttpRequest();
	    
	    req.addEventListener("load", on_load);
	    req.addEventListener("error", on_error);

	    var uri = "https://api.github.com/repos/sfomuseum-data/sfomuseum-data-collection/contents/" + path + "?ref=" + branch;	    
	    req.open("PUT", uri, true);

	    req.setRequestHeader("Accept", "application/vnd.github.v3.json");
	    req.setRequestHeader("Authorization", "token " + _access_token);
	    req.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

	    req.send(form_data);    
	}
	
    };

    return self;

})();    
