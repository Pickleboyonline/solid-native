import SNSwiftUI
import SNCore
import SwiftUI
import YogaSwiftUI

struct ContentView: View {
    @StateObject private var hostReceiver = ObservableHostReceiver()
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
                // Render the root node from the host receiver
                if let rootWrapper = hostReceiver.receiver.getRootNode() {
                    rootWrapper.render()
                        .frame(maxWidth: .infinity, maxHeight: .infinity)
                } else {
                    Text("No root view")
                }
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
                print("Initializing SolidNative Core...")

                // Initialize the Rust core with our delegate
                try hostReceiver.receiver.initializeCore()
                print("Core initialized")

                // Create a root node
                print("Creating root node...")
                let rootId = hostReceiver.receiver.createRoot(tag: "sn_view")
                print("Root node created with ID: \(rootId)")

                // Run JavaScript that builds the UI
                // For now, we'll use inline JS. Later this will load from the dev server.
                print("Building UI with JavaScript...")
                let jsCode = """
                // Get the root view
                const rootId = solidNative.getRootView();

                // Set root container styles
                solidNative.setProp(rootId, 'style', {
                    flex: 1,
                    alignItems: 'center',
                    justifyContent: 'center',
                    backgroundColor: '#ffffff'
                });

                // Create title text
                const titleId = solidNative.createElement('sn_text');
                solidNative.setProp(titleId, 'text', 'Hello Solid Native!');
                solidNative.setProp(titleId, 'style', {
                    fontSize: 24,
                    fontWeight: 'bold',
                    color: '#333333'
                });
                solidNative.insertBefore(rootId, titleId);

                // Create spacer
                const spacer1Id = solidNative.createElement('sn_view');
                solidNative.setProp(spacer1Id, 'style', { height: 20 });
                solidNative.insertBefore(rootId, spacer1Id);

                // Create image
                const imageId = solidNative.createElement('sn_image');
                solidNative.setProp(imageId, 'source', { uri: 'https://picsum.photos/200' });
                solidNative.setProp(imageId, 'style', { width: 200, height: 200 });
                solidNative.setProp(imageId, 'resizeMode', 'cover');
                solidNative.insertBefore(rootId, imageId);

                // Create spacer
                const spacer2Id = solidNative.createElement('sn_view');
                solidNative.setProp(spacer2Id, 'style', { height: 20 });
                solidNative.insertBefore(rootId, spacer2Id);

                // Create subtitle text
                const subtitleId = solidNative.createElement('sn_text');
                solidNative.setProp(subtitleId, 'text', 'Built with Rust + SolidJS + SwiftUI');
                solidNative.setProp(subtitleId, 'style', {
                    fontSize: 14,
                    color: '#666666'
                });
                solidNative.insertBefore(rootId, subtitleId);

                'UI built successfully';
                """

                let result = try hostReceiver.receiver.evalJs(jsCode)
                print("JavaScript result: \(result)")

                // Update UI on main thread
                DispatchQueue.main.async {
                    isLoading = false
                    hostReceiver.objectWillChange.send()
                    print("UI rendering complete!")
                }

            } catch let error as SncoreError {
                print("SNCore Error: \(error)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = "SNCore Error: \(error.localizedDescription)"
                }
            } catch {
                print("Error: \(error)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = error.localizedDescription
                }
            }
        }
    }
}

/// Observable wrapper around HostReceiver for SwiftUI state management
class ObservableHostReceiver: ObservableObject {
    let receiver = HostReceiver()
}
