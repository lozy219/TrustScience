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
