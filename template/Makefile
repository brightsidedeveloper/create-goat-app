.PHONY: build dev


build-wasm:
	cd wasm && GOOS=js GOARCH=wasm go build -o ../public/main.wasm . && cd .. & \
	touch src/main.js

dev:
	(npm start & \
	find wasm -name "*.templ" | entr -r templ generate & \
	find wasm -name "*.go" | entr -r make build-wasm & \
	wait) || (kill 0; exit 1)