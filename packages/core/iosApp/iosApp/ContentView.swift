import SwiftUI
import shared

SolidNative.shared.setFunction {
    6
}

struct ContentView: View {
    let greet = Greeting().greet()


	var body: some View {
		Text(greet)
	}
}

struct ContentView_Previews: PreviewProvider {
	static var previews: some View {
		ContentView()
	}
}
