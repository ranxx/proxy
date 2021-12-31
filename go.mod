module github.com/ranxx/proxy

go 1.17

require (
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.1
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.42.0
)

require (
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)

require (
	github.com/ranxx/observer v0.0.0-20211229065005-3bdd5f9deb1f
	github.com/ranxx/ztcp v0.0.0-20211229061011-7a59f9f991c0
	github.com/spf13/cobra v1.3.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/ranxx/ztcp => ../ztcp
// replace github.com/ranxx/observer => ../observer
