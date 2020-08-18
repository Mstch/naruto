export GOOS=linux
export GOARCH=amd64
/usr/local/go/bin/go build -o ./app github.com/Mstch/naruto #gosetup
docker build -t raft .
docker-compose up -d --scale raft=4