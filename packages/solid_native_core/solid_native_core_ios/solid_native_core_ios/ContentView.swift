import SwiftUI
import shared

func pizza () {
    
}

struct ContentView: View {
	let greet = Greeting().greet()

	var body: some View {
        Text(greet).onAppear() {
            SolidNative.shared.setFunction {
                KotlinInt(10)
            }
        }
        Button("Click Me!") {
            let block = SolidNative.shared.getFunction();
            
            if (block != nil) {
                print(block?() ?? "No value!")
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
	static var previews: some View {
		ContentView()
	}
}
