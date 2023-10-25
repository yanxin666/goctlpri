build:
	go build -ldflags="-s -w" goctlpri.go
	$(if $(shell command -v upx), upx goctlpri)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o goctlpri-darwin goctlpri.go
	$(if $(shell command -v upx), upx goctlpri-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o goctlpri.exe goctlpri.go
	$(if $(shell command -v upx), upx goctlpri.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o goctlpri-linux goctlpri.go
	$(if $(shell command -v upx), upx goctlpri-linux)

image:
	docker build --rm --platform linux/amd64 -t kevinwan/goctlpri:$(version) .
	docker tag kevinwan/goctlpri:$(version) kevinwan/goctlpri:latest
	docker push kevinwan/goctlpri:$(version)
	docker push kevinwan/goctlpri:latest
	docker build --rm --platform linux/arm64 -t kevinwan/goctlpri:$(version)-arm64 .
	docker tag kevinwan/goctlpri:$(version)-arm64 kevinwan/goctlpri:latest-arm64
	docker push kevinwan/goctlpri:$(version)-arm64
	docker push kevinwan/goctlpri:latest-arm64
