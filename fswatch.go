package main

import(
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"os/exec"
	"flag"
	"os"
	"log"
	"os/signal"
	"strings"
	"path/filepath"
	"syscall"
)

// Adds all sub directories of the given directory to the watcher
// and returns the number of watched directories.
func watchRecursive(dir string, watcher *fsnotify.Watcher) int {
	var watched int
	entries, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Println(err)
		return 0
	}

	watcher.Watch(dir)
	watched++

	for _, e := range entries {
		name := e.Name()

		if e.IsDir() && !strings.HasPrefix(name, ".") {
			watched += watchRecursive(filepath.Join(dir, name), watcher)
		}
	}

	return watched
}

func startWatchLoop(watcher *fsnotify.Watcher, quit chan bool) {
	go func() {
		sem := make(chan int, 1)

		for {
			select {
			case <-watcher.Event:
				sem <- 1

				go func() {
					argv := flag.Args()
					
					proc := exec.Command(argv[0], argv[1:]...)
					proc.Stderr = os.Stderr
					proc.Stdout = os.Stdout

					log.Printf("Running command %q\n", proc.Path)

					if err := proc.Run(); err != nil {
						log.Printf("Command %q failed: %s\n", proc.Path, err)
					}

					<-sem
				}()
			case err := <-watcher.Error:
				log.Println(err)
			case <-quit:
				return
			}
		}
	}()
}

var (
	dir string
)

func init() {
	cwd, _ := os.Getwd()
	flag.StringVar(&dir, "dir", cwd, "Directory to watch")

	flag.Parse()
}

func main() {
	rlimit := new(syscall.Rlimit)
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, rlimit)

	rlimit.Cur = rlimit.Max

	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, rlimit)

	if err != nil {
		log.Panicf("Could not change Rlimit: %q", err)
	}

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatalf("Could not create watcher: %s", err)
	}

	quitWatcher := make(chan bool)

	startWatchLoop(watcher, quitWatcher)
	watched := watchRecursive(dir, watcher)

	log.Printf("Watcher started in %q", dir)
	log.Printf("Found %d entries to watch.", watched)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-interrupt:
		watcher.Close()

		log.Println("Shutting down")
		quitWatcher <- true
		log.Println("Done")
	}
}

