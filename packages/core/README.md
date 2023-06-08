TODO:
- Get JS code to Kotlin
- Export type definition file for associated native modules.
- Create VDOM system based off of React Native Fabric implementation (Might not need to be as complex)
- Create webserver that can compile and bundle ts => js code
- Create Kotlin module to download code and run it, also it must expose native APIs
- Build custom solid renderer that can interface with Kotlin VDOM
- Create RN-esk components based off of Swift UI and Jetpack compose
- Copy Expo Modules like API to create Kotlin multiplatform modules.
- Get Chrome debugging to work
- Refactor into library

Main issue deals with types, this is mainly because Swift needs wrapping to send primative data
tp KMM with objc

How Data Commuication Works:
- JS to Native:
    - Module defines basic types. The types are a wrapper over what is essentially a
     "any" type, that gets cast and type checked by specifiying a primative value.
     The actual communcation is a boxed typed and depending on the generic at runtime
     used, the type is validated.
     Module defines functions akin to Expo module APi.
    - Module defines functions that JS can call
- Native to JS:
    - Done by using event handlers and by sending events
    - Events must be registured.

- Views:
    - Views are presented in module as well. When Solid JSX encorters it,
    it only references it by name. However, when its registered it will look up
    the module to create a view with that node.

- Essentially, recreate the example module with the expo modules reference.
  All things considered, solid native is just a collection of these modules,
  with an interface to manage views for solid js to call and bindings for native
  methods.