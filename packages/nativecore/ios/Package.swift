// swift-tools-version:5.3
import PackageDescription


let package = Package(
    name: "Snmobile",
    platforms: [
        .macOS(.v10_14), .iOS(.v13), .tvOS(.v13)
    ],
    products: [
        // Products define the executables and libraries a package produces, and make them visible to other packages.
        .library(
            name: "Snmobile",
            targets: ["Snmobile"])
    ],
    dependencies: [
        // Dependencies declare other packages that this package depends on.
        .package(url: "https://github.com/facebook/yoga.git", from: "3.0.0")
    ],
    targets: [
        // Targets are the basic building blocks of a package. A target can define a module or a test suite.
        // Targets can depend on other targets in this package, and on products in packages this package depends on.
        // Run `make build-mobile-pkg`
        .binaryTarget(
            name: "Snmobile",
            path: "./../build/ios/Snmobile.xcframework"
        )

    ]
)