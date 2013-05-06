package main

import "flag"
import "fmt"
import "log"
import "path"
import "os"

import curl "github.com/edmonds/golang-curl"

var (
    flagVerbose = flag.Bool("v", false, "verbose mode")
)

func main() {
    defer curl.GlobalCleanup()

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [<OPTION>...] <URL>...\n", os.Args[0])
        flag.PrintDefaults()
    }
    flag.Parse()

    easy, err := curl.EasyInit()
    if err != nil {
        log.Fatalln(err)
    }
    defer easy.Cleanup()

    if err := easy.SetOpt(curl.OPT_VERBOSE, *flagVerbose); err != nil {
        log.Fatalln(err)
    }

    urls := flag.Args()
    if len(urls) == 0 {
        flag.Usage()
        os.Exit(1)
    }

    for _, url := range urls {
        fname := path.Base(url)
        _, err := os.Stat(fname)
        if err == nil || (err != nil && !os.IsNotExist(err)) {
            fmt.Fprintf(os.Stderr, "%s: file %s already exists\n", os.Args[0], fname)
            os.Exit(1)
        }

        file, err := os.Create(fname)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s: os.Create(%q) failed: %s\n", os.Args[0], fname, err)
            os.Exit(1)
        }
        defer file.Close()

        if err := easy.SetOpt(curl.OPT_URL, url); err != nil {
            log.Fatalln(err)
        }

        if err := easy.SetOptWrite(curl.IOWriterCallback, file); err != nil {
            log.Fatalln(err)
        }

        if err := easy.Perform(); err != nil {
            log.Fatalln(err)
        }
    }
}
