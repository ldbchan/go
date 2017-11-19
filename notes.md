# Special Features in Go
Here listed my notes on learning Go, mostly about the main differences from other languages.

## Switch
- switch only runs the selected rules, so the `break` statement is not needed
- switch cases can be non-constants (e.g. variables)
- switch values can be non-integers

```go
iPhone := "iPhone"
device := iPhone

switch device {
case "Apple Watch":
    // won't be called
    fmt.Println("Apple Watch")
case iPhone:
    fmt.Println("I got an iPhone")
default:
    // won't be called
    fmt.Println("Are you an Android user?")
}
```

- switch without a condition can be used as a clean way to replace if-else chains

```go
t := time.Now()

switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17: 
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

## Defer, Panic, Recover
It’s a mechanism kind of like try-catch in other languages. ([Defer, Panic, and Recover - The Go Blog](https://blog.golang.org/defer-panic-and-recover))

### defer
- deferred function calls are pushed onto a stack. when a function returns, its deferred calls are executed in last-in-first-out order
- deferred functions may read and assign to the returning function's named return values
```go
// this function prints "3210"
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}

// this function returns 2
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

### panic
- panic stops the ordinary flow of control and begins panicking
- when the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic
- the process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes

### recover
- recover regains control of a panicking goroutine
- recover is only useful inside deferred functions
- during normal execution, a call to recover will return nil and have no other effect
- if the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution

### how defer, panic, and recover work together 

```go
// This program will output:
//
// Calling g.
// Printing in g 0
// Printing in g 1
// Printing in g 2
// Printing in g 3
// Panicking!
// Defer in g 3
// Defer in g 2
// Defer in g 1
// Defer in g 0
// Recovered in f 4
// Returned normally from f.

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0) // g(0) calls panic, f() stops here
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

