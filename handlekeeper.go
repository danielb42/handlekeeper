package handlekeeper

import (
	"log"
	"os"
	"path"

	"github.com/tywkeene/go-fsevents"
)

var (
	InputFile *os.File
)

func OpenFile(file string) error {
	var err error
	InputFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	go inotifyListener(file)
	return err
}

func Close() error {
	return InputFile.Close()
}

func inotifyListener(file string) {
	options := &fsevents.WatcherOptions{
		Recursive: false,
	}

	w, err := fsevents.NewWatcher(path.Dir(file), fsevents.FileRemovedEvent, options)
	if err != nil {
		log.Fatal(err)
	}

	w.StartAll()
	go w.Watch()

	for {
		event := <-w.Events

		if event.IsFileRemoved() && event.Path == file {
			OpenFile(file)
		}
	}
}
