window.addEventListener("load", function load(event){

    console.log("HELLO");
    
    console.log("WORLD");
    
    var map = geotag.maps.getMapById("map");

    if (! map){
	console.log("Failed to retrieve map");
	return;
    }
    
    var layers_control = new L.Control.Layers({
	catalog: sfomuseum.maps.catalog,
    });

    map.addControl(layers_control);
});
