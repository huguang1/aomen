var f = document.getElementById('myform');
//输入框获取
var a = document.getElementById('account');
//提交表单的事件监听
f.onsubmit = function (){
    //判断是否存在内容
    if(a.value==''){

        alert('内容为空，请填写用户名');
        //阻止提交表单
        return false;
    }
};







