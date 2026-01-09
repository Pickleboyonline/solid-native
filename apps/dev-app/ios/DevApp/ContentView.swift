import SNSwiftUI
import SNCore
import SwiftUI

struct ContentView: View {
    @StateObject private var manager = SNSwiftUIManager()
    @State private var errorMessage: String?
    @State private var isLoading = true

    var body: some View {
        VStack(spacing: 20) {
            if isLoading {
                ProgressView("Initializing...")
                    .padding()
            } else if let error = errorMessage {
                VStack(spacing: 12) {
                    Image(systemName: "exclamationmark.triangle.fill")
                        .font(.system(size: 48))
                        .foregroundColor(.red)

                    Text("Error")
                        .font(.headline)

                    Text(error)
                        .font(.caption)
                        .foregroundColor(.secondary)
                        .multilineTextAlignment(.center)
                        .padding(.horizontal)

                    Button("Retry") {
                        isLoading = true
                        errorMessage = nil
                        initializeCoreAndRunJS()
                    }
                    .buttonStyle(.borderedProminent)
                }
                .padding()
            } else {
                SNSwiftUIRenderer(viewTree: manager.viewTree)
                    .frame(maxWidth: .infinity, maxHeight: .infinity)
            }
        }
        .onAppear {
            initializeCoreAndRunJS()
        }
    }

    private func initializeCoreAndRunJS() {
        // Run on background thread to avoid blocking UI
        DispatchQueue.global(qos: .userInitiated).async {
            do {
                print("🚀 Initializing SolidNative Core...")

                // Initialize the Rust core with our delegate
                try manager.initializeWithCore()
                print("✅ Core initialized")

                // IMPORTANT: Create a root node first!
                print("🌳 Creating root node...")
                let rootId = try manager.createRootFromCore(tag: "vstack")
                print("✅ Root node created with ID: \(rootId)")

                // Run JavaScript that builds UI using the root node
                print("📝 Building UI with JavaScript...")
                let jsCode = """
                // Use the root node we created
                const rootId = '\(rootId)';
                // console.log('Root ID:', rootId);

                // Configure root container
                solidNative.setProperty(rootId, 'spacing', '16');
                solidNative.setProperty(rootId, 'padding', '20');

                // Create a title
                const titleId = solidNative.createTextNode('Hello from SolidNative! 🚀');
                solidNative.setProperty(titleId, 'color', 'blue');
                solidNative.insertNode(rootId, titleId);

                // Create a button
                // const buttonId = solidNative.createElement('button');
                // solidNative.setProperty(buttonId, 'backgroundColor', 'green');
                // solidNative.setProperty(buttonId, 'cornerRadius', '8');
                // solidNative.setProperty(buttonId, 'padding', '12');

                // Create button text
                // const buttonTextId = solidNative.createTextNode('Click Me!');
                // solidNative.setProperty(buttonTextId, 'color', 'white');

                // Assemble button
                // solidNative.insertNode(buttonId, buttonTextId);
                // solidNative.insertNode(rootId, buttonId);

                // Create description
                // const descId = solidNative.createTextNode('This UI was built with Rust + JavaScript!');
                // solidNative.setProperty(descId, 'color', 'gray');
                // solidNative.insertNode(rootId, descId);

                // Return success message
                'UI built successfully';
                """

                let result = try manager.evaluateJavaScript(jsCode)
                print("✅ JavaScript result: \(result)")

                // Print tree structure for debugging
                print("\n📊 View Tree Structure:")
                manager.debugPrintTree()

                // Update UI on main thread
                DispatchQueue.main.async {
                    isLoading = false
                    print("✅ UI rendering complete!")
                }

            } catch let error as SnCoreError {
                print("❌ SNCore Error: \(error)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = "SNCore Error: \(error.localizedDescription)"
                }
            } catch {
                print("❌ Error: \(error)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = error.localizedDescription
                }
            }
        }
    }
}
