# handlekeeper

`handlekeeper` is a wrapper for `os.OpenFile()`. It intends to help applications with keeping track of active textfiles by presenting "stable" file handles even when the opened files are moved or deleted. This is achieved by instantly creating a new, empty file in the location of the original file, and re-opening the corresponding filehandle internally.

## Usage / Example
Here, `/var/log/myApp.log` can be moved or deleted without having to reopen file handles in the reading/writing application. 

```
hk := handlekeeper.NewHandlekeeper("/var/log/myApp.log")
defer hk.Close()
for {
	scanner := bufio.NewScanner(hk.Handle)
	for scanner.Scan() {
		println(scanner.Text())
	}
	
	time.Sleep(time.Second)
	}
}

```
