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