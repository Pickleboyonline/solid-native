package com.example.solid_native_core

import platform.Foundation.NSNumber
import platform.Foundation.NSString
import platform.JavaScriptCore.JSContext
import platform.JavaScriptCore.JSExportProtocol
import platform.JavaScriptCore.JSValue
import platform.JavaScriptCore.setObject
import platform.UIKit.UIDevice
import platform.darwin.NSObject


class Calculator {
    fun add(x: Int, y: Int) = x + y
}

val c = Calculator()

val jsContext = JSContext();

@Suppress("CAST_NEVER_SUCCEEDS")
fun runner(): String {
    jsContext.exceptionHandler = { _, value: JSValue? ->
        if (value != null) {
            println("Error: $value")
        } else {
            println("JS Error!")
        }
    }

    // But in function block
    jsContext.setObject({ x: NSNumber, y: NSNumber ->
            val result = (x as Double) + (y as Double)
            return@setObject result as NSNumber
        },
        "add" as NSString
    );

    jsContext.setObject({ x: Double, y: Double ->
        x + y
    },
        "addTwoNumbers" as NSString
    );

    println("40 + 3: " + jsContext.evaluateScript("addTwoNumbers(40,3)").toString())

    // Call function and evaluate;
    return jsContext.evaluateScript("add(0,33223)").toString()
}

class IOSPlatform: Platform {
    override val name: String = runner()
     // UIDevice.currentDevice.systemName() + " " + UIDevice.currentDevice.systemVersion
}

actual fun getPlatform(): Platform = IOSPlatform()


//final class LL : JSExportProtocol {
//
//}