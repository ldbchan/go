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

