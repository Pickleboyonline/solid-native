package com.example.solid_native_core

var function: (() -> Int)? = null

object SolidNative {

    fun setFunction(block: () -> Int) {
        function = block
    }

    fun getFunction() = function

    fun registerModule() {

    }
}