# solid-native
 React Native but for Solid.JS

React Native is nice, but it has the following issues:
1. Difficult upgrades and installations
2. Difficulty creating native modules and interfacing with native modules easily
3. Almost impossible debugging with new hermes engine
4. Poor monorepo support

This project will most likely borrow much from the new architecture that react native 
has proposed (some research required here). I want to replace React.JS with Solid JS and
the C++ new architecture rendering unit with Kotlin multi-platform mobile.
(Rust is also an option, but Kotlin has an advantage since its already supports debugging,
and some enhanced cross platform native capabilities.)

To address difficult upgrades and installations, I am to make the main package 1 dependency, 
with 1 package, and 1 ios script. If additional build scripts are needed, I aim to make them
as minimal as possible.

I want to make this as modern and future-proof as possible, so I will opt for deno to manage 
the codebase.

Deno has good monorepo support, considering there is no concept of "packages", just files.

Look into what expo has done with their module system to create a multi module system.

Also look at what NativeScript has done to generate automatic bindings into native modules.

There are some interesting articles on how SwiftUI and Jetpack Compose have been integrated with
react native, which would prove useful.



