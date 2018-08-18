BINARY=spider

VERSION=1.0.0

BUILD=`date +%FT%T%z`

# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.Env=production -s -w"



default:
	go build -o ${BINARY} -v ${DEV_LDFLAGS} -tags=jsoniter

install:
	govendor sync -v