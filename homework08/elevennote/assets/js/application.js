require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {
    $("#note-Title").attr({
        "placeholder": "Title your note",
    });
    $("#note-Tag").attr({
        "data-role": "tagsinput"
    });
});
