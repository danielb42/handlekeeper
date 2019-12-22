package handlekeeper

import (
	"os"
	"path"

	fse "github.com/tywkeene/go-fsevents"
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
	w, err := fse.NewWatcher()
	if err != nil {
		return err
	}

	_, err = w.AddDescriptor(path.Dir(file), fse.FileRemovedEvent)
	if err != nil {
		return err
	}

	if err = w.StartAll(); err != nil {
		return err
	}

	go w.Watch()

	go func() {
		defer w.StopAll() //nolint:errcheck

		for {
			event := <-w.Events

			if event.IsFileRemoved() && event.Path == file {
				if err = hk.openFile(file); err != nil {
					break
				}
			}
		}
	}()

	return nil
}
