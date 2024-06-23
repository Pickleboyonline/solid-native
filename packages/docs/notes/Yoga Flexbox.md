This task is done, below are just some notes.

---
TODO: Essentially copy the way react fabric system works under the hood.
Important Links:
- [https://reactnative.dev/architecture/fabric-renderer](https://reactnative.dev/architecture/fabric-renderer)
- [https://www.yogalayout.dev/](https://www.yogalayout.dev/)
- [https://github.com/software-mansion-labs/react-native-swiftui/tree/master/packages/react-native-swiftui/native/Views](https://github.com/software-mansion-labs/react-native-swiftui/tree/master/packages/react-native-swiftui/native/Views)
- [https://github.com/software-mansion-labs/react-native-swiftui/tree/master](https://github.com/software-mansion-labs/react-native-swiftui/tree/master)
- [https://blog.swmansion.com/swiftui-renderer-for-react-native-2b62fda38c9b](https://blog.swmansion.com/swiftui-renderer-for-react-native-2b62fda38c9b)
- [https://blog.swmansion.com/swiftui-renderer-for-react-native-part-ii-161bad1d69e1](https://blog.swmansion.com/swiftui-renderer-for-react-native-part-ii-161bad1d69e1)

Essentially, when a style is placed on a node we call yoga to have a node for it. This may be a bit tricky considering SolidJS idology.

SolidJS works via mutations to the tree to update. But to render something, we need the information up front about the styles (one pass) on initial mount. Otherwise, we run into the problem of a view being present but no object is sent down. 

I have not seen this visibly, but i am just saying it can cause problems.

React is better in that it knows the tree up front in every update the the VDOM. technically, my implimentation needs a VDOM of its own to manage the references to the views to update the props. However, as long as the system does it in one pass (SolidJS may group updates, im not exactly sure, because when i view somethings typically dont shift they render almost instantly.)

WAIT never mind SolidJS builds the updates before attatching the views anyways. That makes sense. OK, so we just need to call to YOGA when we MOUNT or UPDATE a PROP on a dirty component.

Because we are working with solid, we may be able to get some selectively out of this. In the render phase.

In this new layout system, every component is in a ZStack with its offset and frame ajusted to YOGA dimensions. so Yoga needs to calculate params. Where does this happen? Well, it can happen in the RENDER module. The render module is a singleton that dispatches commands to the components based on IDs.

So, the renderer will just need to maintain a yoga tree based on its methods. If a node has not been attachted, it will skip calculating layouts (i think yoga automatically does this, so your good.)

The only thing to worry about is batching updates after the initial mount. If the update happens via the styles object thats not bad since its just one object but if its outside of that than a performance hit can be observed. That can be worked on later, but for now, its not that bad and probably wont be that noticeable.

So the data flow looks like this:

SolidJS => SolidNativeRenderer => Yoga(yoga may call some native methods for text and device dimensions) => A layout object => SolidNativeView which uses its layout for flexbox.

In the SolidNativeRenderer, there needs to be something that recalculates the layout after each prop change if the component has mounted. If it hasn't mounted, wait until all its props have been updated (so really just calculate it when you draw it)

The code for determining the size of text or views can be derrived from ChatGPT and the Software mansion blog.

In terms of threading, all this is done on the main for now but we can move the solid native renderer prob on the background when performing Yoga layouts.

I wanted to avoid using UIKit because apparantly you cant use it for widgets or apple watch, but thats not that bad. If i do use pure swiftUI, some parts are just not that feasible like rendering offscreen components. You can *technically* do with with a hidden zstack but that means having to render things twice on the main thread with some trick sync issues. Overall, because SwiftUI is meant to be used declaratively means I have to determine everything upfront including the layout.

I could also just re-impliment the Yoga flexbox with SwiftUI native primatives but thats too much work in my opinion. ChatGPT probably wont be able to do something that complex so I cant rely on that either. Best option is just just repurpose Yoga, which should take a day or too, and i get the flexbox system for free basically.

