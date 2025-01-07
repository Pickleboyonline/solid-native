// swift-tools-version:5.3
import PackageDescription

let package = Package(
  name: "Snmobile",
  platforms: [
    // .macOS(.v10_14), .iOS(.v13),
    .iOS(.v13),
    .macOS(.v10_14),
  ],
  products: [
    .library(
      name: "Snmobile",
      targets: ["SnmobileWrapper"]
    )
    // .library(name: "QuickJS", targets: ["CQuickJS"]),
  ],
  dependencies: [
    .package(url: "https://github.com/facebook/yoga.git", from: "3.1.0")
  ],
  targets: [
    .binaryTarget(
      name: "Snmobile",
      path: "./../build/ios/Snmobile.xcframework"
    ),
    .target(
      name: "QuickJS",
      path: "Sources/QuickJS",
      sources: ["src"],
      publicHeadersPath: "include",
      cSettings: [
        .define("CONFIG_VERSION", to: "\"2023-12-09\""),
        .define("_GNU_SOURCE"),
        .define("CONFIG_BIGNUM"),
        .headerSearchPath("include"),
      ]
    ),
    .target(
      name: "SnmobileWrapper",
      dependencies: [
        .target(name: "Snmobile"),
        .product(name: "yoga", package: "yoga"),
        .target(name: "QuickJS"),
      ]
      // cSettings: [
      //   .headerSearchPath("vendor/quickjs")
      // ],
    ),
  ]
)
