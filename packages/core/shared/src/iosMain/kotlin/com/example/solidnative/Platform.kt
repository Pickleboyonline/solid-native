package com.example.solidnative

import platform.JavaScriptCore.JSContext
import platform.UIKit.UIDevice


val context = JSContext()

class IOSPlatform: Platform {
    override val name: String = context.evaluateScript("1 + 22").toString()
//UIDevice.currentDevice.systemName() + " " + UIDevice.currentDevice.systemVersion
}

actual fun getPlatform(): Platform = IOSPlatform()