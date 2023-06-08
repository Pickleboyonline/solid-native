package com.example.solidnative

import kotlin.native.concurrent.ThreadLocal


var function: (() -> Int)? = null

object SolidNative {

    fun setFunction(block: () -> Int) {
        function = block
    }

    fun getFunction() = function
}