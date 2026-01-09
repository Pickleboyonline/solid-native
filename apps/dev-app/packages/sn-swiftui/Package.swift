// swift-tools-version:5.9
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "SNSwiftUI",
    platforms: [
        .iOS(.v15),
        .macOS(.v12)
    ],
    products: [
        .library(
            name: "SNSwiftUI",
            targets: ["SNSwiftUI"]
        )
    ],
    dependencies: [
        // Local dependency on the generated SNCore package
        .package(path: "../sncore/build/swift/SNCore")
    ],
    targets: [
        .target(
            name: "SNSwiftUI",
            dependencies: ["SNCore"],
            path: "Sources/SNSwiftUI"
        ),
        .testTarget(
            name: "SNSwiftUITests",
            dependencies: ["SNSwiftUI"],
            path: "Tests/SNSwiftUITests"
        )
    ]
)
