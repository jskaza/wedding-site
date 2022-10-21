$(document).ready(function () {

    var dropdown = document.getElementById('displayName');
    var response = document.getElementById('response');
    var yes = document.getElementById('yes');
    var no = document.getElementById('no');
    var review = document.getElementById('review');

    dropdown.addEventListener(
        'change', (event) => {
            response.hidden = false;
            review.hidden = false;
        })

    $('#rsvp').change(function () {
        selectedValue = $("input[name='rsvpOptions']:checked").val();
        if (selectedValue == "no") {
            no.hidden = false
            yes.hidden = true
        } else if (selectedValue == "yes") {
            no.hidden = true
            yes.hidden = false
        }
    })


})