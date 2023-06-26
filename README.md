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

## Current Example Showcase

https://github.com/Pickleboyonline/solid-native/assets/16438041/2bb1975b-24bf-4863-89cb-563aa2961116

## How to run:

Make sure you have Deno installed. Then, run the following command to start the bundler:
```ts
deno task start
```
The bundler should then be running on port 8080.

From there, open the iOS app in `packages/ios`. Run the app in the simulator and it should load the test app in `packages/test_app`. Note, live reload/hot reload is not supported yet, so re-run the iOS app when you make a change to the test app!

## Update/Future Roadmap

After investigating NativeScript, it seems to feature most of what I want in a cross platform mobile development. It has easy access to Native APIs with auto codegen and typescript supports. Project upgrades look to be easier since it generates the Android and iOS project code as needed depending on the NativeScript version. No more cumbersome upgrades in the Xcode/Android Studio Project like React Native! Another nice feature is that you put native code in the `App_Resources` folder and these files are automatically liked into the native projects and typescript bindings are generated for them. No more complex TurboModules setup and having to rewrite APIs!

However, the documentation is a bit lackluster compared to that of React Native. I want this project to be an extension of NativeScript with a focus on SolidJS support and wrapping SwiftUI and JetPack Compose components. As a result, the documentation for this project should be easier to manage since I only focus on supporting the mentioned frameworks.

Some other changes I want to make:
- Deno Support
- Webpack => Esbuild

I did some research on how to ahead of time compile the app but I could not find a compiler. It turns out Javascript/Typescript is <i>really</i> hard to statically compile if not impossible in most cases. I investigate whether the (hopc)[http://hop.inria.fr/home/index.html] compiler could generate faster code than using some interpreters, but it was significantly slower. I computed a fib sequence up to 30 with O(n^2) complexity and hopc took around 1 minute to complete whereas the JS interpreters quickjs, hermes, v8jitless all took around 10 seconds. With jit mode enabled for V8, the program took around 1 second to run which is about the same time the equivalent C++ program I wrote took.

So, the only way I can realistically see a way to optimize JS into native code is the following:
- Statically analyze the JS code for deterministic behavior where the types can be known at compile time
- Split that code and transpile it into a C++ function, then put a reference to that C++ function in the javascript code
- Convert the Javascript code into bytecode for the JS Interpreter and bind the C++ function into the global context
- Convert into an executable/compile

The issue is that yes in theory this creates more performant code, but the question is how much JS code can we optimize at compile time
in order to achieve significant results? This method is somewhat of what hopc is supposed to do but for some reason it could not optimize the fib function despite me passing in 2 numbers; therefore, the compiler could have known that function only needs a number signature and produce an optimized version of it.

Why do all of this work? Well JIT is still supreme for JS execution but in some environments you either can't run with JIT (iOS) or there aren't enough resources to enable it (low power embedded devices that have limited memory and space). An AOT compiler could offer potential speed ups for iOS JS powered apps.

To accomplish this task, I would need to known how to manipulate TS/JS and C++ Abstract Syntax Trees. Once I have a good grasp of how the language is constructed and how to transform/transpile it, then I can build algorithms that could do static typing analysis to see if any native code can be generated from JS.

## TODO

- Import my custom solid-js render into a NativeScript project.
- Determine how to integrate my SwiftUI elements into the root view
    - I think NativeScript uses UIKit but my solid_native project does not. To make this easier I need to create some sort of UIKit wrapper to get into my SwiftUI components mounted in.
- Get Webpack to transpile SolidJS JSX. Also need to reconfigure global typings for my components.
- Continue creating SwiftUI components.
    - This should be a lot easier since the bindings can be autogenerated.
- Make documentation
- Potentially fork nativescript and get Deno support working along with ESBuild 