<!DOCTYPE>
<html>
<head>
{{template "header.tpl"}}
</head>
    <body>
        <div class="whole-width">
            <div class="login-part">
                <h1>登陆</h1>
                <form action="/login" method="post">
                    <ul>
                        <li>
                            <label>用&emsp;户&emsp;名:</label><input type="text" name="username">
                        </li>
                        <li>
                            <label>密&emsp;&emsp;&emsp;码:</label><input type="password" name="pwd">
                        </li>
                        <li class="form-operate">
                           <a href="/regist" style="width:80px;">注册</a>
                           <button type="submit">登录</button>
                        </li>
                    </ul>
                </form>
            </div>
        </div>
    </body>
</html>
