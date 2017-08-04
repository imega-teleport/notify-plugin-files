CWD=/go/src/github.com/imega-teleport/notify-plugin-files

test: build
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) \
		golang:1.8-alpine sh -c "go list ./... | grep -v 'vendor\|integration' | xargs go test"
	@docker run -d --name nginx_stub -v $(CURDIR)/tmp:/data -p 80:80 imega/nginx-stub
	@docker run --rm -v $(CURDIR)/rel:/usr/local/bin -v $(CURDIR):/data -w /data --link nginx_stub:server alpine sh -c 'sender -url "http://server:80" -storageUrl storage.ru -user authuser -pass authpass -path "tests/fixtures"'
	@docker run --rm -v $(CURDIR)/tests:/tests -v $(CURDIR)/tmp:/data alpine sh -c 'diff /tests/expected-response.txt /data/*'

build: rel/sender

rel/sender:
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=0 \
		golang:1.8-alpine \
		sh -c 'apk add --update --no-cache git && \
		go build -v -o $@'

clean:
	@-docker rm -fv nginx_stub
	@-rm -rf $(CURDIR)/tmp
