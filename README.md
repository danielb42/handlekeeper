# handlekeeper

Wrapper to os.OpenFile() to avoid moving file handles.

When a file is moved (e.g. rotated) open file handles travel along with it. 
`handlekeeper` instantly reopens a file handle to a newly created file in the original location.

## Example
```

func main() {
    handlekeeper.OpenFile("/var/log/rotating.log")
    defer handlekeeper.Close()
    
    scanner := bufio.NewScanner(handlekeeper.InputFile)
    ...

    # oh no, now logrotate moved "rotating.log" to "rotating.log.1".
    # normally my scanner would have travelled along, but it 
    # continues to read from "rotating.log". thanks, handlekeeper!
}
```
