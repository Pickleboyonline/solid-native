package com.example.solidnative

import platform.JavaScriptCore.JSContext
import platform.JavaScriptCore.JSExportProtocol




val context = JSContext();



class IOSPlatform: Platform {
    override val name: String = context.evaluateScript("1 + 22").toString()
//UIDevice.currentDevice.systemName() + " " + UIDevice.currentDevice.systemVersion
}

actual fun getPlatform(): Platform = IOSPlatform()

// Calls linked Swift libraries and registures there defnitions
fun registerModule() {}



//
