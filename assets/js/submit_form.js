function formJSON() {
    let obj = $('form').serializeJSON();
    var jsonString = JSON.stringify(obj);


    $.ajax({
        type: "POST",
        url: "/rsvp/guests",
        data: jsonString,
        success: function (result) {
            console.log(result);
        },
        dataType: "json"
    });

}
