# TrustScience

http://bbs.ngacn.cc/read.php?tid=14044587


## Quick setup guide
- Set up image processing service
```
$ go run backend/main.go
```

- Configure endpoints for local project
  - Open `frontend/main.js`
  - Toggle line 24 and 25
  - You may need to edit `backend/main.go:24` to configure the allow origins header as well.
  
- Start the project with any local http server
```
$ http-server .
```

[CuteLilTwo优化准备]
1.搭建对应数据库
2.页面数据从后端进行SQL查询后展示
3.进行胜率计算前,添加区服名称选择
4.添加自动增加不同区服场次的功能(数据库操作)
5.不同区服的不同场次进行一次记录,进入数据库(数据库操作,根据输入的区服 记录时间进行判断)
6.计算胜率放入后台自动计算
7.后续加入御魂判断
