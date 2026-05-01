; Set the graphics mode (Width, Height, Depth, Mode)
Graphics 640,480,32,2
SetBuffer BackBuffer() ; Use double-buffering for smooth animation

x = 0

While Not KeyHit(1) ; Loop until the Escape key (code 1) is pressed
    Cls ; Clear the screen
    
    Rect x, 100, 50, 50, 1 ; Draw a solid rectangle at (x, 100)
    
    x = x + 1
    If x > 640 Then x = -50 ; Reset position when it goes off-screen
    
    Flip ; Flip the back buffer to the front
Wend

End
