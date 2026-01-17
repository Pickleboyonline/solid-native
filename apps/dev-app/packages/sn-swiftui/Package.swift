// swift-tools-version:5.9
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "SNSwiftUI",
    platforms: [
        .iOS(.v16),
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "SNSwiftUI",
            targets: ["SNSwiftUI"]
        )
    ],
    dependencies: [
        // Local dependency on the generated SNCore package
        .package(path: "../sncore/build/swift/SNCore"),
        // Local dependency on Yoga-SwiftUI for flexbox layout
        .package(path: "../Yoga-SwiftUI"),
        // Yoga layout engine (needed for YG types)
        .package(url: "https://github.com/facebook/yoga.git", from: "3.1.0"),
    ],
    targets: [
        .target(
            name: "SNSwiftUI",
            dependencies: [
                "SNCore",
                .product(name: "YogaSwiftUI", package: "Yoga-SwiftUI"),
                .product(name: "yoga", package: "yoga")
            ],
            path: "Sources/SNSwiftUI"
        ),
        .testTarget(
            name: "SNSwiftUITests",
            dependencies: ["SNSwiftUI"],
            path: "Tests/SNSwiftUITests"
        )
    ]
)
