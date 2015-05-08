$(window).load(function() {
    document.getElementById("uploadBtn").addEventListener("change", handleFileSelect, true);
    document.getElementById("uploadBtn").addEventListener("click",
    function() {
        this.value = null
    },
    true);
});

function handleFileSelect(n) {
    for (var u = n.target.files,
    t, r, i = 0; t = u[i]; i++) t.type.match("image.*") && (r = new FileReader, r.onload = function(n) {
        return function(t) {
            updateThumbnail(t.target.result);
        }
    } (t), r.readAsDataURL(t))
}
function updateThumbnail(n) {
    var t = document.getElementById("thumbnail");
    t.setAttribute("src", n)
}
