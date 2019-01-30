package handlekeeper

import (
	"os"
	"path"

	fse "github.com/tywkeene/go-fsevents"
)

// Handlekeeper holds the actual file handle
type Handlekeeper struct {
	Handle *os.File
}

// NewHandlekeeper returns a *Handlekeeper wrapping a file handle
func NewHandlekeeper(file string) (*Handlekeeper, error) {
	hk := &Handlekeeper{}

	if err := hk.openFile(file); err != nil {
		return nil, err
	}

	if err := hk.startInotifyListener(file); err != nil {
		return nil, err
	}

	return hk, nil
}

func (hk *Handlekeeper) openFile(file string) error {
	fh, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)

	if err == nil {
		hk.Handle = fh
	}

	return err
}

// Close closes the file handle.
func (hk *Handlekeeper) Close() error {
	return hk.Handle.Close()
}

func (hk *Handlekeeper) startInotifyListener(file string) error {
	options := &fse.WatcherOptions{
		Recursive: false,
	}

	w, err := fse.NewWatcher(path.Dir(file), fse.FileRemovedEvent, options)

	if err != nil {
		return err
	}

	w.StartAll()
	go w.Watch()

	go func() {
		defer w.StopAll()

		for {
			event := <-w.Events

			if event.IsFileRemoved() && event.Path == file {
				if err = hk.openFile(file); err != nil {
					break
				}
			}
		}
	}()

	return err
}
