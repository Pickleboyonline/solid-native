OK, so I want the project to start with a DENO cli command, like:

```bash
deno run -A https:://solid-native.com/setup-app.ts
```

Then, sets up similar project to react native and native script. Native files are generated based on typescript source and configuration files.

When one wants to add native code, use some cli command, and a summodule will be created as a Swift and android package, with some typescript code. The files are linked into the core project by the build system.

Modules will the the ios, android, src, and go folders. Go code is added for cross platform code. (not sure if i want this, however, because of the build system. I was considering Kotlin multiplatform (which now that i think about it, MAY have been a better choice considering the Jetpack support) ), but didn't want to initially because the build system is VERY complex in compared to go, which is just a single command. Biggest issue with Go mobile is that the runtime could be duplicated with each go code, which is like ~1.8 MB. Zig/Rust could have been better but: Rust is very complex and increases dev time and Zig is not as production ready & doesnt have a good bindgen system. 

NEVER MIND, gomobile bind can link all into one package. You just have to specifiy mutplie packages. So it DOES work. Example:

```bash
gomobile bind -target ios multipkgs/pkg/a multipkgs/pkg/b
```

So in the build step youd just have to specify all the packages (and Swift Source package system must refere to the same target. Im 99% sure the SPM will just allow it all to reference the same thing.) Well, go is still better than kotlin multiplatform.

Generally, the less you need to touch the native code, the better. Ideally, we want to make it such that one can configure the entire project.

Config will be a file that exports a `configire` function and take in the project object where you call methods on to configure it, like adding modules and such. 

Also, it would be interesting to build the app along with the SolidStart framework and use tRPC with it.
