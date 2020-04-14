window.addEventListener("load", function load(event){

    var map = geotag.maps.getMapById("map");

    var layers_control = new L.Control.Layers({
	catalog: sfomuseum.maps.catalog,
    });

    map.addControl(layers_control);
});
