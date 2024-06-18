build-ios:
	@# TODO: make build iOS script and ensure to go get correct binaries for go mobile
	@# Command to ensure necessary deps: go get -d golang.org/x/mobile/cmd/gomobile
	@echo hello world!

clean-build:
	rm -rf ./build

build: clean-build
	go build -o build/bin/main main.go

build-yoga:
	@echo "Cleaning Prev Yoga Build..."
	cd extenal/yoga
	rm -rf build
	mkdir build
	cd build
	@echo "Building Yoga TODO: Get release flags..."
	# Using arm64 build, change in make for your os!
	cmake -DCMAKE_OSX_ARCHITECTURES=arm64 -DBUILD_SHARED_LIBS=OFF ..
	make