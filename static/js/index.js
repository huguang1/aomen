$(function () {
    var username = cookie.get("username");
    if ("" === username || "undefined" === typeof username) {
        window.location.href = "/static/templates/login.html";
    }
    $("#nameh3").html(username);
    $("#logout").attr("href","/logout?name="+username);
});