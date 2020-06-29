<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>【Union ID】邮箱验证</title>
</head>
<body>
<!-- 邮箱验证 -->
<div style="border-top: 1px dashed #000;border-bottom: 1px dashed #000;">
    <p><br></p>
    <p>{{.UserName}}：</p>
    <p>您好，您本次操作的邮箱验证码为：{{.Code}}</p>
    <p><br></p>
    <p>Dear {{.UserName}}:</p>
    <p>Your email verification code is: {{.Code}}</p>
    <p><br></p>
    <p>此验证码将在邮件发出30分钟后失效。</p>
    <p>It will become invalid in 30 min after this mail is sent successfully.</p>
    <p><br></p>
    <p>如果该请求非本人操作，您的账号可能存在安全风险，建议修改您的密码。</p>
    <p>If this request is not done on your own, security risk may exist in your account. It is recommended to modify your password.</p>
    <p><br></p>
    <p style="text-align: right;">统信软件</p>
    <p style="text-align: right;">Uniontech Software</p>
</div>
</body>
</html>