# handlekeeper

Wrapper for os.OpenFile() - keeps a filehandle pointed to a files' original location even if the file is moved (e.g. rotated) somewhere else or deleted. A new, empty file is created at the location instantly and the known file handle is preserved. 

In a regular scenario the filehandle would move along with the file, thus ceasing to read/write the intended location. `handlekeeper` intends to help applications with keeping track of active textfiles by presenting "stable" file handles.

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
