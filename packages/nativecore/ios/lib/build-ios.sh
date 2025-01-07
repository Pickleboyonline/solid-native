#!/bin/bash

# Directory setup
QUICKJS_DIR="quickjs"
BUILD_DIR="build-ios"
IOS_MIN_VERSION="11.0"

# Download QuickJS if not present
if [ ! -d "$QUICKJS_DIR" ]; then
    git clone https://github.com/bellard/quickjs.git "$QUICKJS_DIR"
fi

# Create build directory
mkdir -p "$BUILD_DIR"

# iOS SDK path
SDK_PATH=$(xcrun --sdk iphoneos --show-sdk-path)

# Common compiler flags
COMMON_FLAGS="-arch arm64 \
    -isysroot $SDK_PATH \
    -mios-version-min=$IOS_MIN_VERSION \
    -fembed-bitcode \
    -fPIC"

# Configure the build
cat > "$QUICKJS_DIR/config.h" << EOF
#define CONFIG_VERSION "2023-12-09"
#define CONFIG_BIGNUM 1
#define CONFIG_LTO 0
EOF

# Modify Makefile for iOS
cd "$QUICKJS_DIR"
cat > Makefile.ios << EOF
PREFIX=/usr/local
CC=clang
AR=ar
STRIP=strip
CFLAGS=$COMMON_FLAGS -O2 -Wall -D_GNU_SOURCE -DCONFIG_VERSION="\\"2023-12-09\\""
LDFLAGS=$COMMON_FLAGS
LIBS=-lm
DEFINES:=-DCONFIG_BIGNUM
EOF

# Build object files
clang $COMMON_FLAGS -c \
    quickjs.c \
    libbf.c \
    libregexp.c \
    libunicode.c \
    cutils.c \
    -I. \
    -O2 -Wall -D_GNU_SOURCE -DCONFIG_VERSION=\"2023-12-09\" \
    -DCONFIG_BIGNUM

# Create static library
ar rcs "../$BUILD_DIR/libquickjs.a" \
    quickjs.o \
    libbf.o \
    libregexp.o \
    libunicode.o \
    cutils.o

# Clean up
rm *.o

echo "Build complete. Static library is in $BUILD_DIR/libquickjs.a"