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
	@echo "Building Yoga for host tests"
	bash external/yoga/unit_tests