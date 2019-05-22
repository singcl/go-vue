# AIM

shadowsocks + proxifier = 直接更新 go 依赖

## hero-cli 部署 go

**Go Dependencies via dep**

https://devcenter.heroku.com/articles/go-apps-with-dep

**HOW to use heroku git**

https://devcenter.heroku.com/articles/git

### 问题：

```sh
# 查看heroku日志
heroku logs --tail
```

heroku 启动不成功，查看日志报错：
`Error R10 (Boot timeout) -> Web process failed to bind to $PORT within 60 seconds of launch`

### 解决方法

heroku 会为 web 应用动态分配 PORT，所以不要固定 PORT 即可

```go
  // r.Run(":3000")
  r.Run()
```

### 示例地址：https://go-vue-example.herokuapp.com/
