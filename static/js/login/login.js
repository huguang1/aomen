var box = {
    content: ''
    , btn: ['确定']
    , yes: function () {
        window.location.href = '/static/templates/login.html';
    }
};

// 用户登陆
layui.use('layer', function () {
    var layer = layui.layer;
    $('#login_button').click(function () {
        var layer = layui.layer;
        var username = $('#username').val();
        var password = $('#password').val();
            $.ajax({
                url: '/config/login',
                type: 'post',
                dataType: 'json',
                data: {
                    'username': username,
                    'password': password,
                },
                success: function (data) {

                    if (data.status === 200) {
                        window.location.href = '/static/view/index.html';
                    } else {
                        box.content = data.message;
                        layer.open(box)
                    }
                },
                error: function () {
                    box.content = '用户名或密码错误';
                    layer.open(box)
                }
            })
    });
});


