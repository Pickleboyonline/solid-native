package com.example.solid_native_core

interface Platform {
    val name: String
}

expect fun getPlatform(): Platform