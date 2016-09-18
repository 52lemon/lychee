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
                            <label>用户名:</label><input type="text" name="username">
                        </li>
                        <li>
                            <label>密&nbsp;&nbsp;码:</label><input type="password" name="pwd">
                        </li>
                        <li class="form-operate">
                           <button type="submit">登录</button>
                        </li>
                        <li>
                           <a href="/regist">注册</a>
                        </li>
                    </ul>
                </form>
            </div>
        </div>
    </body>
</html>
