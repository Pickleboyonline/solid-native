package com.example.solidnative

interface Platform {
    val name: String
}

expect fun getPlatform(): Platform