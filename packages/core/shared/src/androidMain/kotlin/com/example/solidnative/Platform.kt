package com.example.solidnative

import android.content.Context
import androidx.javascriptengine.JavaScriptSandbox
import androidx.startup.Initializer
import com.google.common.util.concurrent.ListenableFuture



class JSSandBoxInitializer : Initializer<ListenableFuture<JavaScriptSandbox>> {
    override fun create(context: Context): ListenableFuture<JavaScriptSandbox> {
        TODO("Not yet implemented")
        val jsSandboxFuture = JavaScriptSandbox.createConnectedInstanceAsync(context);
        return jsSandboxFuture
    }



}


class AndroidPlatform : Platform {
    override val name: String = "Android ${android.os.Build.VERSION.SDK_INT}"
}

actual fun getPlatform(): Platform = AndroidPlatform()