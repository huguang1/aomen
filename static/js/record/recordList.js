layui.use(['table', 'layer', 'upload'], function () {
    var $ = layui.jquery, layer = layui.layer, upload = layui.upload;
    var table = layui.table;
    var tableIns = table.render({
        elem: '#recordList',
        url: '/config/record',
        method: 'get',
        cols: [[{
            field: 'id',
            title: 'ID',
            sort: true,
            align: 'center'
        }, {
            field: 'account',
            title: '会员账号',
            align: 'center'
        }, {
            field: 'month_amount',
            title: '投注金额',
            align: 'center'
        }, {
            field: 'compute',
            title: '状态',
            templet: '#switchState',
            unresize: true}
        , {
            field: 'date',
            title: '日期',
            align: 'center'
        }, {
            title: '操作',
            width: 250,
            align: 'center',
            fixed: 'right',
            toolbar: '#barDemo'
        }]],
        page: true //是否显示分页
        , parseData: function (res) {
            return {
                "code": 0,
                'msg': '',
                "count": res.count,
                "data": res.results
            }
        }
        , limit: 10,
        limits: [5, 10, 100]
        //添加权限控制
    });

    $('#selectbtn').on('click', function () {
        active.reload();
    });
    //监听工具条
    table.on('tool(demo)', function (obj) {
        var data = obj.data;
        if (obj.event === 'del') {
            layer.confirm('真的删除行么', function (index) {
                $.ajax({
                    type: "post",
                    url: "/config/deleterecord",
                    data: {"id": data.id},
                    dataType: "json",
                    success: function (data) {
                        if (data.status === 200) {
                            layer.msg('删除成功');
                            setTimeout(function () {
                                active.reload();
                            }, 1000);
                        }
                    }
                });
                layer.close(index);
            });
        } else if (obj.event === 'edit') {
            var index = layer.open({
                type: 2,
                content: '/static/view/record/editRecord.html?id=' + data.id,
                area: ['65%', '65%'],
                maxmin: true,
                success: function (layero, index) {
                    var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                    var iframeWin = window[layero.find('iframe')[0]['name']];
                    iframeWin.inputMemberInfo(data);
                },
                end: function () {
                    active.reload();
                }
            });
        }
    });
    active = {
        reload: function () {
            var searchAccount = $('#searchAccount');
            // 执行重载
            table.reload('recordList', {
                page: {
                    curr: 1
                },
                where: {
                    search_account: searchAccount.val()
                }
            });
        }
    };

    //执行实例
    var uploadInst = upload.render({
        elem: '#test3' //绑定元素
        , url: '/upload' //上传接口
        , accept: 'file'
        , data: {}
        , done: function (res) {
            layer.msg(res.message);
            active.reload();
        }
        , error: function () {

        }
    });

    // 添加
    $('#addbtn').on('click', function () {
        layer.open({
            type: 2,
            title: false,
            content: '/static/view/record/addRecord.html',
            area: ['65%', '65%']
        });
    });
});
