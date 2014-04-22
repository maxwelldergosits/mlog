package mlog

import (
  "io"
  "log"
  "os"
  "io/ioutil"
  "path/filepath"
  "os/user"
  "strings"
  "time"
)

func getFileWriter(prefix string,disableLog bool) io.Writer {

  if disableLog {
    logFile := ioutil.Discard
    return logFile
  }

  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  if (prefix == "") {
    prefix = dir
  }
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }

  homeDir := usr.HomeDir
  if prefix[:1] == "~" {
    prefix = strings.Replace(prefix, "~", homeDir, 1)
  }

  var logFile io.Writer

  var time = time.Now()

  const RFC3339 = "2006-01-02T15:04:05Z07:00"
  var logFileName = prefix+"/"+time.Format(RFC3339)+".log"

  log.Println("log file:",logFileName)
  logFile,err = os.Create(logFileName)
  if err != nil {
    log.Fatal(err)
  }

  return logFile

}
