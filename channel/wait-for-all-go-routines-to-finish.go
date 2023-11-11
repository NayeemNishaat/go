func main() {
    c := make(chan struct{}) // We don't need any data to be passed, so use an empty struct
    for i := 0; i < 100; i++ {
        go func() {
            doSomething()
            c <- struct{}{} // signal that the routine has completed
        }()
    }

    // Since we spawned 100 routines, receive 100 messages.
    for i := 0; i < 100; i++ {
        <- c
    }
}
