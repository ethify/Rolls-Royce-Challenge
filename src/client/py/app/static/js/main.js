$(function(){
    $('.content-box-large').on("click","#btn-group1", function(){
        $.ajax({
            url: '/group1',
            data: $('#form-group').serialize(),
            type: 'POST',
            success: function(){
            $("#alert-g1").show();
            $("#alert-g1").fadeOut(3000);
            //$('.content-box-large').load(document.URL +  ' .content-box-large');
            },
            error: function(error){
                console.log(error);
            }
        });
    });
});

$(function(){
    $('.content-box-large').on("click","#btn-group2", function(){
        $.ajax({
            url: '/group2',
            data: $('#form-group2').serialize(),
            type: 'POST',
            success: function(){
            $("#alert-g1").show();
            $("#alert-g1").fadeOut(3000);
            //$('.content-box-large').load(document.URL +  ' .content-box-large');
            },
            error: function(error){
                console.log(error);
            }
        });
    });
});

$(function(){
    $('.content-box-large').on("click","#btn-group3", function(){
        $.ajax({
            url: '/admin',
            data: $('#form-group3').serialize(),
            type: 'POST',
            success: function(){
            $("#alert-g1").show();
            $("#alert-g1").fadeOut(3000);
            //$('.content-box-large').load(document.URL +  ' .content-box-large');
            },
            error: function(error){
                console.log(error);
            }
        });
    });
});

$(function(){
    $('#polls-row').on("click", '.vote-btn',function(){
        $.ajax({
            url: '/vote',
            data: $(this).parent().serialize(),
            type: 'POST',
        }).done(function(response){
            console.log(response);
            //location.reload(true);
            $('#polls-row').load(document.URL +  ' #polls-row');
        });
    });
});
