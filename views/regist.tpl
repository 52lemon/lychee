<!DOCTYPE>
<html>
    <head>
    {{template "header.tpl"}}
    </head>
    <body>
        <div class="whole-width">
            <div class="login-part">
                <h1>注册</h1>
                <form action="/regist" method="post">
                <ul>
                    <li>
                        <label>用&emsp;户&emsp;名:</label><input type="text" name="username">
                    </li>
                    <li>
                        <label>密&emsp;&emsp;&emsp;码:</label><input type="password" name="pwd">
                    </li>
                    <li class="form-operate">
                        <a href="/login" style="width:80px;">登陆</a>
                        <button type="submit">注册</button>
                    </li>
               </form>
            </div>
        </div>
    </body>
</html>
