module music-pc-server

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/pprof v1.2.1
	github.com/gin-gonic/gin v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.2.0
	github.com/smallnest/rpcx v0.0.0-20200104074109-4f758db059eb
	github.com/valyala/fasthttp v1.8.0
	github.com/xinliangnote/go-util v0.0.0-20191116000206-e64f4ad6c381
	go.uber.org/zap v1.13.0
	google.golang.org/grpc v1.24.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
