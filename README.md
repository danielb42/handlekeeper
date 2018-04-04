# handlekeeper

Wrapper to os.OpenFile() to avoid moving file handles.

When a file is moved (e.g. rotated) open file handles travel along with it. 
`handlekeeper` instantly reopens a file handle to a newly created file in the original location.

## Example
```

func main() {
    err := handlekeeper.OpenFile("/var/log/rotating.log")
    defer handlekeeper.Close()

    # oh no, logrotate moved rotating.log to rotating.log.1
    
    scanner := bufio.NewScanner(handlekeeper.InputFile)

    # nice, that scanner is still aimed at "rotating.log" ...
}
```
