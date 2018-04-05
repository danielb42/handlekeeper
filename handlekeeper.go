package handlekeeper

import (
	"log"
	"os"
	"path"

	"github.com/tywkeene/go-fsevents"
)

type Handlekeeper struct {
	Handle *os.File
}

func NewHandlekeeper(file string) (*Handlekeeper, error) {
	hk := &Handlekeeper{}
	err := hk.openFile(file)
	hk.startInotifyListener(file)

	return hk, err
}

func (hk *Handlekeeper) openFile(file string) error {
	fh, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	hk.Handle = fh
	return nil
}

func (hk *Handlekeeper) Close() error {
	return hk.Handle.Close()
}

func (hk *Handlekeeper) startInotifyListener(file string) {
	options := &fsevents.WatcherOptions{
		Recursive: false,
	}

	w, err := fsevents.NewWatcher(path.Dir(file), fsevents.FileRemovedEvent, options)
	if err != nil {
		log.Fatal(err)
	}

	w.StartAll()
	go w.Watch()

	go func() {
		defer w.StopAll()

		for {
			event := <-w.Events

			if event.IsFileRemoved() && event.Path == file {
				err := hk.openFile(file)
				if err != nil {
					break
				}
			}
		}
	}()
}
