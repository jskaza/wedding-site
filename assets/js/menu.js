$(document).ready(function () {
    $("[href]").each(function () {
        if (this.href == window.location.href) {
            $(this).parents('span').fadeOut();
            $(this).addClass("active");
        }
    })
})