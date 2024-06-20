build-mobile-pkg: clean-build
	@# Command to ensure necessary deps:
	@go get -d golang.org/x/mobile/cmd/gomobile
	@mkdir -p build/android
	@mkdir -p build/ios
	@echo "Building iOS and Android Packages..." 
	@gomobile bind -target=android -androidapi 34 -o build/android/snmobile.aar nativecore/lib/snmobile
	@gomobile bind -target=ios,iossimulator -prefix SN -o build/ios/Snmobile.xcframework nativecore/lib/snmobile


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