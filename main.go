package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
  flag "github.com/ogier/pflag"
  "bytes"
)

const path string = "/home/enju/.config/zenbu/variable_sets/"

var help bool
var list bool

func init() {
  flag.BoolVarP(&help, "help", "h", false, "Display this help message")
  flag.BoolVarP(&list, "list", "l", false, "List the templates")
  flag.Parse()
}

func main() {
  if help == true {
    PrintHelpMessage()
  }
  if list == true {
    listTemplates()
  }
  arg := os.Args[1]
  if arg != "" {
    fmt.Println("this")
    set()
  }
}

func set() {
  arg := os.Args[1]
  files, _ := readNames()
  var theme string
  var set bool
  for _, file := range files {
    if file == arg {
      theme = file
      set = true
    }
  }
  if set == true {
    fmt.Println(theme)
  }
}

func listTemplates() {
  files, _ := readNames()
  for _, base := range files {
    if filepath.Ext(base) != "" {
      fmt.Println(strings.TrimSuffix(base, filepath.Ext(base)))
    }
  }
  os.Exit(1)
}

func readNames() ([]string, error) {
  searchDir := "/home/enju/.config/zenbu/variable_sets/"

  fileList := []string{}
  err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
    base := filepath.Base(path)
    var bf bytes.Buffer
    if filepath.Ext(base) == "" {
      cFolder := strings.TrimSuffix(base, filepath.Ext(base))
      bf.WriteString(cFolder)
      bf.WriteString("/")
    } else {
      base = strings.TrimSuffix(base, filepath.Ext(base))
      bf.WriteString(base)
    }

    fmt.Println(bf.String())

    fileList = append(fileList, bf.String())
    return nil
  })

  if err != nil {
    return nil, nil
  }

  return fileList, nil
}

func PrintHelpMessage() {
  fmt.Printf("Usage: %s [options]\n", os.Args[0])
  fmt.Printf("Options:\n")
  flag.PrintDefaults()
  os.Exit(1)

}
