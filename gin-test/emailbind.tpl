<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>【Union ID】绑定邮箱验证</title>
</head>
<body>
<!-- 绑定邮箱验证 -->
<div style="border-top: 1px dashed #000;border-bottom: 1px dashed #000;">
    <p>亲爱的用户：</p>
    <p>您好，您本次操作的邮箱验证码为：{{.Code}}</p>
    <p><br></p>
    <p>Dear:</p>
    <p>Your email verification code is: {{.Code}}</p>
    <p><br></p>
    <p>此验证码将在邮件发出30分钟后失效。</p>
    <p>It will become invalid in 30 min after this mail is sent successfully.</p>
    <p><br></p>
    <p>如果该请求非本人操作，请忽略本邮件。</p>
    <p>If this request is not done on your own, please ignore this mail.</p>
    <p><br></p>
    <p style="text-align: right;">统信软件</p>
    <p style="text-align: right;">Uniontech Software</p>
</div>
</body>
</html>