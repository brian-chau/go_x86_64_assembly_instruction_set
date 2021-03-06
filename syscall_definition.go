package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func process_csv_line_by_line( filePath string, column string, value string ) {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file: " + filePath, err)
    }
    defer f.Close()

    r := csv.NewReader(f)
    is_first_line := true
    header := make(map[string]int)
    var results [9]string // One index for each column
    fmt.Println("-----")

    for {
        record, err := r.Read()
        if err == io.EOF {
            return
        }
        if err != nil {
            log.Fatal(err)
        }
        if is_first_line {
            is_first_line = false
            for i := range record {
                header[record[i]] = i
            }
        } else {
            if strings.Contains(record[header[column]], value) {
                for h, i := range header {
                    results[i] = fmt.Sprintf("%-15s: %v",h, record[i])
                }
                for _, i := range results {
                    fmt.Println(i)
                }
                fmt.Println("-----")
            }
        }
    }

    return
}

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Invalid number of arguments.")
    }

    filename := "x86_syscalls.csv"
    //filename := "x86_64_syscalls.csv"
    column   := "syscall_name"
    value    := os.Args[1]
    process_csv_line_by_line( filename, column, value )
}
