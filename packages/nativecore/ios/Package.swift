// swift-tools-version:5.3
import PackageDescription

let package = Package(
  name: "SNLib",
  platforms: [
    // .macOS(.v10_14), .iOS(.v13),
    .iOS(.v13),
    .macOS(.v10_14),
  ],
  products: [
    .library(
      name: "SNLib",
      targets: ["SNLibStub"]
    )
    // .library(name: "QuickJS", targets: ["CQuickJS"]),
  ],
  dependencies: [
    .package(url: "https://github.com/facebook/yoga.git", from: "3.1.0")
  ],
  targets: [
    .binaryTarget(
      name: "SNLib",
      path: "./../build/ios/SNLib.xcframework"
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
      name: "SNLibStub",
      dependencies: [
        .target(name: "SNLib"),
        .product(name: "yoga", package: "yoga"),
        .target(name: "QuickJS"),
      ]
      // cSettings: [
      //   .headerSearchPath("vendor/quickjs")
      // ],
    ),
  ]
)
