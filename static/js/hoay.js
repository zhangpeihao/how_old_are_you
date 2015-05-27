$(window).load(function() {
    document.getElementById("uploadBtn").addEventListener("change", handleFileSelect, true);
    document.getElementById("uploadBtn").addEventListener("click",
    function() {
        this.value = null
    },
    true);

    $("#img_div").width($("#thumbnail").width());
});

function handleFileSelect(n) {
    for (var u = n.target.files,
    t, r, i = 0; t = u[i]; i++) t.type.match("image.*") && (r = new FileReader, r.onload = function(n) {
        return function(t) {
            // updateThumbnail(t.target.result);

            var api = new FacePP('0ef14fa726ce34d820c5a44e57fef470', '4Y9YXOMSDvqu1Ompn9NSpNwWQFHs1hYD');
            api.request('detection/detect', {
                url: 'http://www.faceplusplus.com.cn/wp-content/themes/faceplusplus/assets/img/demo/16.jpg'
                // img: t.target.result
            }, function(err, result) {
            if (err) {
                // TODO handle error
                return;
            }
            // TODO use result
            // document.getElementById('response').innerHTML = JSON.stringify(result, null, 2);
            console.log(result);
            updateThumbnail(result.url);
            for(var i in result.face) {
                var face_width = result.img_width * result.face[i].position.height / 100 + 12;
                var face_height = result.img_height * result.face[i].position.height / 100 + 12;
                var face_left = result.img_width * result.face[i].position.center.x / 100 - face_width / 2;
                var face_top = result.img_height * result.face[i].position.center.y / 100 - face_height / 2;
                var age_left = result.img_width * result.face[i].position.center.x / 100 - 40;
                var age_top = face_top - 60;
                var html = '<div style="position:absolute; z-index:5; top: ' + face_top + 'px; left: ' + face_left + 'px; width: ' + face_width + 'px; height: ' + face_height + 'px; border:6px solid red;"></div><div style="position:absolute; z-index:5; top: ' + age_top + 'px; left: ' + age_left + 'px; width: 80px; height:60px; background:#f1d100;"><p style="font-size:40px; text-align:center;">' + result.face[i].attribute.age.value + '</p></div>';

                $("#img_div").append(html);
            }
            });
        }
    } (t), r.readAsDataURL(t))
}
function updateThumbnail(n) {
    var t = document.getElementById("thumbnail");
    t.setAttribute("src", n)
}
