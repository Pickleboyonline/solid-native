Need to create a `measureNode` function for the text component (and text layout).

Essentially, I need to create a function that receives the props of text. For now, we'll get fontSize.

Ok, now I was able to get the font size of a text component. 

I'll try to replicate what the text view is doing when its contrained like that. Maybe its being clipped?

Not exactly sure what react native does if the text is overflown.

Yoga does know size, but it needs to work on the text itself and its rendering process.

Expected behavior:
Text should clip when the frame is below, similar to how the other system works.

Text DOES clip and respect measurements.
Its really more an issue of UIKit size vs SwiftUI. Need to ensure the styles match up, but otherwise fine.

Done => just need to ensure swiftui fonts are same defaults as UIKits, but otherwise code i wrote is fine.

https://www.yogalayout.dev/docs/advanced/external-layout-systems#invalidating-measurements

^^ Need to mark node as dirty when text changes.

https://reactnative.dev/docs/text#containers

OK, so text nodes DO NOT have associated Yoga nodes. 

How react native text components work:
- Text is bounded to PARENT. Even if it wraps around.
- ![[simulator_screenshot_90729C35-E1AB-47E2-8C5F-EFD0AC87BE03.png]]
- In the above image, the "word" block is bounded to the container width/height, despite having the "hello world" there. Implementation wise, this is quite simple to impliment.

So, basically, heres the jist:
- If its a text node, you can keep the yoga node and regular stuff. Internally its the same. You could remove the yoga node for it, but for now ill just leave it as is. (You can make a special node that references that type, but since we do pointers its different. Youd have to Box it in another struct as a union type. (Go sadly doesnt have rust style enum union structs w/ pattern match :( )))
- Host is NOT nofitfied of any text components that have 
- Yoga nodes of child text components are not notified. 
- 
