dev_env:
	#export GOPATH=$HOME/github.com/rcrespodev/Blogs/design/repository/
	export PATH=/usr/local/go/bin
	export REDIS_HOST=localhost
	export REDIS_PORT=6379
	export VENDOR_ENDPOINT=https://api.coindesk.com/v1/bpi/currentprice.json
	export APP_HOST=localhost
	export APP_PORT=8080

redis_instance:
#	sudo docker container rm redis-test-instance
	sudo docker pull redis
	sudo docker run --name redis-test-instance -p 6379:6379 -d redis

golang_test:
	go test ./...

tests: dev_env redis_instance golang_test
