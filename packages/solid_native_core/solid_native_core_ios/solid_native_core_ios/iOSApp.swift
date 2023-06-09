import SwiftUI
import JavaScriptCore

//class Props: ObservableObject {
//    @Published var message = ""
//    @Published var children: [SNView] = []
//}
//
//struct Component: View {
//    @ObservedObject var props: Props
//
//    var body: some View {
//        Text("Hello, \(props.message)!")
//    }
//}
//
//
//class SNView {
//    let props = Props();
//
//    @ViewBuilder
//    func render() -> some View {
//        Component(props: props)
//    }
//}

class Props: ObservableObject {
    @Published var message = "World"
}

struct Component: View {
    @ObservedObject var props: Props

    var body: some View {
        Text("Hello, \(props.message)!")
    }
}

class JSObjectModule {
    let jsContext: JSContext
    let jsObject: JSValue
    
    init(context: JSContext) {
        self.jsContext = context
        self.jsObject = JSValue.init(newObjectIn: context)
    }
    
    func addSyncFunction(name: String, executor: @escaping @convention(block) () -> Void) {
        // context.setObject(executor, forKeyedSubscript: name as NSString)
        // jsObject.setObject(executor, forKeyedSubscript: "name")
        jsObject.setValue(executor, forProperty: name)
    }
    
    func addProperty(key: String, value: Any) {
        jsObject.setValue(value, forKey: key)
    }
    
}


let globalJsContext = JSContext()
/**
 All you need to do is define a struct that reacts to props,
 mutate the props based on JS events.
 Children are iterated on  with ForEach and calling `render()`
 */
class SNView {
    let props = Props()
    
    // Props object contains everthing we need to know. For children.
    // Mutating that array will cause SwiftUI to re-render
    
    // The classes have a binding to JS
    
    func defineJsReference() -> JSObjectModule {
        let jsModule = JSObjectModule(context: globalJsContext!)
        
        jsModule.addSyncFunction(name: "print") {
            print("Hello World!")
        }
        
        return jsModule
    }
    
    /**
     I can make a more elagant API for this, but this is the jist.
     */
    func setMessage(newMessage: String) {
        props.message = newMessage
    }
    
    
    @ViewBuilder
    func render() -> some View {
        Component(props: props)
    }
}



@main
struct iOSApp: App {
    let snViewBuilder = SNView()
    
    init() {
        // snViewBuilder.world.objectWillChange.send()
        let module = snViewBuilder.defineJsReference();
        
        globalJsContext?.setObject(module.jsObject, forKeyedSubscript: "SNModule" as NSString);
        
        print("Results: " + (globalJsContext?.evaluateScript("SNModule.print()").toString())!)
        
    }
    
	var body: some Scene {
		WindowGroup {
            snViewBuilder.render()
		}
	}
}
