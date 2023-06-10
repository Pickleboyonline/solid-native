import SwiftUI
import JavaScriptCore

// Reference!!!
@objc protocol PersonJSExports: JSExport {
    var firstName: String { get set }
    var lastName: String { get set }
    var birthYear: NSNumber? { get set }

    var fullName: String { get }

    // Imported as `Person.createWithFirstNameLastName(_:_:)`
    static func createWith(firstName: String, lastName: String) -> Person
}

@objc public class Person : NSObject, PersonJSExports {
    // Properties must be declared with `dynamic`
    dynamic var firstName: String
    dynamic var lastName: String
    dynamic var birthYear: NSNumber?

    required init(firstName: String, lastName: String) {
        self.firstName = firstName
        self.lastName = lastName
    }

    var fullName: String {
        return "\(firstName) \(lastName)"
    }

    class func createWith(firstName: String, lastName: String) -> Person {
        return Person(firstName: firstName, lastName: lastName)
    }
}

@objc public class DummyValue: NSObject {
    
}

@main
struct iOSApp: App {
    // let snViewBuilder = SNView()
    let rootElement: SNElementNode
    
    init() {
        let core = SNCore()
        rootElement = core.createElementNode(name: "Root")
        core.registureGlobalObject()
        let context = core.jsContext
        context.setObject(rootElement.getJsObjectExport(), forKeyedSubscript: "rootView" as NSString)
        // rootElement.getJsObjectExport().setProp(name: "name", value: "World")
        print(context.evaluateScript(
        """
        const view = SNCore.createElement("TextView")
        view.isTextNode
        rootView.setProp("name", "world")
        """
        ).toString()!)
    }
    
	var body: some Scene {
		WindowGroup {
            rootElement.value.render()
		}
	}
}
