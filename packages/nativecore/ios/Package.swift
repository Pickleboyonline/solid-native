// swift-tools-version:5.3
import PackageDescription

let package = Package(
  name: "Snmobile",
  platforms: [
    // .macOS(.v10_14), .iOS(.v13),
     .iOS(.v13),
     .macOS(.v10_14)
  ],
  products: [
    .library(
      name: "Snmobile",
      targets: ["SnmobileWrapper"]
    )
  ],
  dependencies: [
    .package(url: "https://github.com/facebook/yoga.git", from: "3.0.0")
  ],
  targets: [
    .binaryTarget(
      name: "Snmobile",
      path: "./../build/ios/Snmobile.xcframework"
    ),
    .target(
      name: "SnmobileWrapper",
      dependencies: [
        .target(name: "Snmobile"),
        .product(name: "yoga", package: "yoga"),
      ]
    ),
  ]
)
