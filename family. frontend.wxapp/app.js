// app.js
App({
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        if (res.code) {
            wx.request({
                method: 'POST',
                url: 'http://127.0.0.1:7070/api/login/wechat_login',
                data: {
                    code: res.code
                },
                success: res => {
                    debugger
                    if (res.statusCode === 200) {
                        console.log(res.data.sessionId)// 服务器回包内容
                    }
                }
            })
        }else {
            console.log('获取用户登录态失败！' + res.errMsg)
        }
      }
    })
  },
  globalData: {
    userInfo: null
  }
})
