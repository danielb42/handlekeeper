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

func NewHandlekeeper(file string) *Handlekeeper {
	hk := &Handlekeeper{}
	hk.openFile(file)
	hk.startInotifyListener(file)

	return hk
}

func (hk *Handlekeeper) openFile(file string) {
	fh, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	hk.Handle = fh
}

func (hk *Handlekeeper) Close() error {
	return hk.Close()
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
		for {
			event := <-w.Events

			if event.IsFileRemoved() && event.Path == file {
				hk.openFile(file)
			}
		}
	}()
}
