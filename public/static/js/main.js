$(document).ready(function(){
	
	var $distanceForm = $("#distanceForm");
	
	$("[type=submit]", $distanceForm).click(function(event){
		event.preventDefault();
		$('#alert, #distanceResult, #distanceResult2, #address1, #address2').hide();
		
		var $from = $("#from", $distanceForm);
		var $to = $("#to", $distanceForm);
		validateRequired($from);
		validateRequired($to);
		info("from: " + $from.val() + " to: " + $to.val());
		if($('.error', $distanceForm).length > 0){
			showFormError();
		}
		
		var distanceUrl = "/ds/distance/" + $from.val() + "/" + $to.val() + "/1";
		$.getJSON(distanceUrl, function(data) {
			showDistanceResult(data);
		});
		distanceUrl = "/ds/distance/" + $from.val() + "/" + $to.val() + "/2";
		$.getJSON(distanceUrl, function(data) {
			showDistanceResult2(data);
		});
		
		var addressUrl = "/ds/address/" + $from.val();
		$.getJSON(addressUrl, function(data) {
			showAddressResult(data);
		});
		addressUrl = "/ds/address/" + $to.val();
		$.getJSON(addressUrl, function(data) {
			showAddressResult2(data);
		});
	});
	
});
function getAddressHtml(data) {
	var addressHtml = 
		"<address>" +
		"<strong>"+data.Street+"</strong><br>" + 
		data.Town + ", " + data.County + "<br>" +
		data.Postcode + "<br>" +
		data.Country +
		"</address>";
	return addressHtml;
}
function showAddressResult(data){
	var addressHtml = getAddressHtml(data);
	$('#address1').html(addressHtml).fadeIn();
}
function showAddressResult2(data){
	var addressHtml = getAddressHtml(data);
	$('#address2').html(addressHtml).fadeIn();
}
function showDistanceResult(data){
	var $d = $('#distanceResult');
	$('.label', $d).text(parseInt(data.Meters) + " meters")
	$d.fadeIn();
}
function showDistanceResult2(data){
	var $d = $('#distanceResult2');
	$('.label', $d).text(parseInt(data.Meters) + " meters")
	$d.fadeIn();
}
function info(msg) {
	$('#info').prepend('<li>' + msg + '</li>');
}

function showFormError(msg) {
	if(!msg)
		msg = 'Oops something is wrong, check errors below and try again';
	$('#alert').text(msg).fadeIn();
}

function validateRequired($elErr) {
	var isValid = !($elErr.val() == null || $elErr.val().trim() == '');
	
	if(isValid) {
		$elErr.next('p').remove();
		$elErr.parents('.control-group').removeClass('error');
	}else{  	
		$elErr.next('p').remove().end().after('<p class="help-block">Field is required.</p>');
		$elErr.parents('.control-group').addClass('error');
	}
}