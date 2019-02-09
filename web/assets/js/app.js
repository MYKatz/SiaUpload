$(document).ready( function() {

    //register fileselect trigger
    $(document).on('change', ':file', function() {
        var input = $(this),
            label = input.val().replace(/\\/g, '/').replace(/.*\//, '');
        input.trigger('fileselect', [label]);
    });
    //handle file select
    $('#fileUpload').on('fileselect', function(event, label) {
        $('#uploadLabel').html(label)
    });


});