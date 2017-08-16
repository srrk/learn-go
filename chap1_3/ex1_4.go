// Read file contents in bulk and process them in bulk
// for duplicate entries and print them.
// ex1.4 - In addition add filenames in which duplicate entries
// are found.
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    filenames := "" // filenames which contain duplicates
    for _, filename := range os.Args[1:] {
        // Intialize a map for every filename.
        file_count := make(map[string]int)
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        // do file wise map count.
        for _, line := range strings.Split(string(data), "\n") {
            file_count[line]++
        }
        // find any duplicate in this filecount.
        for _, n := range file_count {
            if n > 1 {
                filenames += filename
                break
            }
        }
        // append to the aggregate map.
        for line, _ := range file_count {
            counts[line]++
        }

    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
    // Print filenames which have duplicate entries.
    fmt.Printf("Duplicates found in files : %s", filenames)
}
