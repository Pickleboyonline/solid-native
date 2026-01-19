import SNSwiftUI
import SNCore
import SwiftUI
import YogaSwiftUI
import Combine

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
                ScrollView {
                    VStack(spacing: 16) {
                        Image(systemName: "exclamationmark.triangle.fill")
                            .font(.system(size: 48))
                            .foregroundColor(.red)

                        Text("Error")
                            .font(.headline)

                        Text(error)
                            .font(.system(.body, design: .monospaced))
                            .foregroundColor(.primary)
                            .multilineTextAlignment(.leading)
                            .frame(maxWidth: .infinity, alignment: .leading)
                            .padding()
                            .background(Color(.systemGray6))
                            .cornerRadius(8)

                        Button("Retry") {
                            isLoading = true
                            errorMessage = nil
                            initializeCoreAndRunJS()
                        }
                        .buttonStyle(.borderedProminent)
                    }
                    .padding()
                }
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
                solidNative.setProp(imageId, 'source', { uri: 'https://picsum.photos/100' });
                solidNative.setProp(imageId, 'style', { width: 200, height: 100 });
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
    
            } catch let error as SnCoreError {
                let errorDetail = formatSnCoreError(error)
                print("SNCore Error: \(errorDetail)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = errorDetail
                }
            } catch {
                print("Error: \(error)")
                DispatchQueue.main.async {
                    isLoading = false
                    errorMessage = "Unknown Error:\n\(error)"
                }
            }
        }
    }
    }

//    private func initializeCoreAndRunJS() {
//        let devServerURL = URL(string: "http://localhost:8080/bundle.js")!
//
//        Task { @MainActor in
//            do {
//                try hostReceiver.receiver.initializeCore()
//                _ = hostReceiver.receiver.createRoot(tag: "sn_view")
//
//                let (data, response) = try await URLSession.shared.data(from: devServerURL)
//
//                guard let httpResponse = response as? HTTPURLResponse else {
//                    throw NSError(domain: "DevServer", code: -1, userInfo: [NSLocalizedDescriptionKey: "Invalid response from dev server"])
//                }
//
//                guard httpResponse.statusCode == 200 else {
//                    throw NSError(domain: "DevServer", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Dev server returned status \(httpResponse.statusCode)"])
//                }
//
//                guard let jsCode = String(data: data, encoding: .utf8) else {
//                    throw NSError(domain: "DevServer", code: -2, userInfo: [NSLocalizedDescriptionKey: "Could not decode JS bundle as UTF-8"])
//                }
//
//                _ = try hostReceiver.receiver.evalJs(jsCode)
//
//                isLoading = false
//                hostReceiver.objectWillChange.send()
//
//            } catch let error as SnCoreError {
//                isLoading = false
//                errorMessage = formatSnCoreError(error)
//            } catch {
//                isLoading = false
//                errorMessage = "Failed to load JS bundle:\n\(error.localizedDescription)\n\nMake sure the dev server is running:\ndeno task dev"
//            }
//        }
//    }
//}

/// Observable wrapper around HostReceiver for SwiftUI state management
@MainActor
final class ObservableHostReceiver: ObservableObject, @unchecked Sendable {
    let receiver = HostReceiver()
}

/// Formats SnCoreError with detailed information
func formatSnCoreError(_ error: SnCoreError) -> String {
    switch error {
    case .RuntimeError(let msg):
        return "Runtime Error:\n\(msg)"
    case .RendererError(let msg):
        return "Renderer Error:\n\(msg)"
    case .JsEvalError(let msg):
        return "JavaScript Evaluation Error:\n\(msg)"
    case .ContextError(let msg):
        return "Context Error:\n\(msg)"
    }
}

// Imperative, no solidjs rendering:
//private func initializeCoreAndRunJS() {
//    // Run on background thread to avoid blocking UI
//    DispatchQueue.global(qos: .userInitiated).async {
//        do {
//            print("Initializing SolidNative Core...")
//
//            // Initialize the Rust core with our delegate
//            try hostReceiver.receiver.initializeCore()
//            print("Core initialized")
//
//            // Create a root node
//            print("Creating root node...")
//            let rootId = hostReceiver.receiver.createRoot(tag: "sn_view")
//            print("Root node created with ID: \(rootId)")
//
//            // Run JavaScript that builds the UI
//            // For now, we'll use inline JS. Later this will load from the dev server.
//            print("Building UI with JavaScript...")
//            let jsCode = """
//            // Get the root view
//            const rootId = solidNative.getRootView();
//
//            // Set root container styles
//            solidNative.setProp(rootId, 'style', {
//                flex: 1,
//                alignItems: 'center',
//                justifyContent: 'center',
//                backgroundColor: '#ffffff'
//            });
//
//            // Create title text
//            const titleId = solidNative.createElement('sn_text');
//            solidNative.setProp(titleId, 'text', 'Hello Solid Native!');
//            solidNative.setProp(titleId, 'style', {
//                fontSize: 24,
//                fontWeight: 'bold',
//                color: '#333333'
//            });
//            solidNative.insertBefore(rootId, titleId);
//
//            // Create spacer
//            const spacer1Id = solidNative.createElement('sn_view');
//            solidNative.setProp(spacer1Id, 'style', { height: 20 });
//            solidNative.insertBefore(rootId, spacer1Id);
//
//            // Create image
//            const imageId = solidNative.createElement('sn_image');
//            solidNative.setProp(imageId, 'source', { uri: 'https://picsum.photos/200' });
//            solidNative.setProp(imageId, 'style', { width: 200, height: 200 });
//            solidNative.setProp(imageId, 'resizeMode', 'cover');
//            solidNative.insertBefore(rootId, imageId);
//
//            // Create spacer
//            const spacer2Id = solidNative.createElement('sn_view');
//            solidNative.setProp(spacer2Id, 'style', { height: 20 });
//            solidNative.insertBefore(rootId, spacer2Id);
//
//            // Create subtitle text
//            const subtitleId = solidNative.createElement('sn_text');
//            solidNative.setProp(subtitleId, 'text', 'Built with Rust + SolidJS + SwiftUI');
//            solidNative.setProp(subtitleId, 'style', {
//                fontSize: 14,
//                color: '#666666'
//            });
//            solidNative.insertBefore(rootId, subtitleId);
//
//            'UI built successfully';
//            """
//
//            let result = try hostReceiver.receiver.evalJs(jsCode)
//            print("JavaScript result: \(result)")
//
//            // Update UI on main thread
//            DispatchQueue.main.async {
//                isLoading = false
//                hostReceiver.objectWillChange.send()
//                print("UI rendering complete!")
//            }
//
//        } catch let error as SnCoreError {
//            let errorDetail = formatSnCoreError(error)
//            print("SNCore Error: \(errorDetail)")
//            DispatchQueue.main.async {
//                isLoading = false
//                errorMessage = errorDetail
//            }
//        } catch {
//            print("Error: \(error)")
//            DispatchQueue.main.async {
//                isLoading = false
//                errorMessage = "Unknown Error:\n\(error)"
//            }
//        }
//    }
//}
//}
