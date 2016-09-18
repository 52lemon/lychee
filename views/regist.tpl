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
                        <label>用户名:</label><input type="text" name="username">
                    </li>
                    <li>
                        <label>密&nbsp;&nbsp;码:</label><input type="password" name="pwd">
                    </li>
                    <li class="form-operate">
                        <button type="submit">注册</button>
                    </li>
                    <li>
                        <a href="/login">登陆</a>
               </form>
            </div>
        </div>
    </body>
</html>
