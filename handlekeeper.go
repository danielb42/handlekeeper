package handlekeeper

import (
	"os"
	"path"

	"github.com/tywkeene/go-fsevents"
)

type Handlekeeper struct {
	Handle *os.File
}

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

func (hk *Handlekeeper) Close() error {
	return hk.Handle.Close()
}

func (hk *Handlekeeper) startInotifyListener(file string) error {
	options := &fsevents.WatcherOptions{
		Recursive: false,
	}

	w, err := fsevents.NewWatcher(path.Dir(file), fsevents.FileRemovedEvent, options)

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
