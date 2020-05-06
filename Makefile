
PROJECT := suggest-wasm
REGISTRY := gcr.io/no-proj-yet

TARGET_WASM := web/main.wasm
TARGET_WASM_GZ := web/main.wasm.gz

LATEST := $(PROJECT):latest
TEST_PORT := 8088

NGINX_MOUNTS := -v$(shell pwd)/ngninx.conf:/etc/nginx/nginx.conf \
	-v$(shell pwd)/web:/usr/share/nginx/html:ro


all: $(TARGET_WASM) $(TARGET_WASM_GZ) web/wasm_exec.js


$(TARGET_WASM): main.go
	GOOS=js GOARCH=wasm go build -o $(TARGET_WASM)

$(TARGET_WASM_GZ): $(TARGET_WASM)
	gzip -k $(TARGET_WASM)

web/wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./web/

image: all
	docker build -t $(LATEST) -t $(REGISTRY)/$(LATEST) .

serve_mounted: all
	# now visit: http://localhost:$(TEST_PORT)/
	docker run --name $(PROJECT) --rm -p$(TEST_PORT):80 $(NGINX_MOUNTS) nginx

serve_image: image
	# now visit: http://localhost:$(TEST_PORT)/
	docker run --name $(PROJECT) --rm -p$(TEST_PORT):80 $(LATEST)

image_push:
	docker push $(REGISTRY)/$(LATEST)

clean:
	rm -f $(TARGET_WASM) $(TARGET_WASM_GZ) web/wasm_exec.js
