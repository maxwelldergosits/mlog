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


  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  if (prefix == "") {
    prefix = dir
  }
  usr, _ := user.Current()
  homeDir := usr.HomeDir
  if prefix[:1] == "~" {
    prefix = strings.Replace(prefix, "~", homeDir, 1)
  }

  var logFile io.Writer

  var time = time.Now()

  const RFC3339 = "2006-01-02T15:04:05Z07:00"
  var logFileName = prefix+"/"+time.Format(RFC3339)+".log"

  if disableLog {

    logFile = ioutil.Discard
  } else {
    log.Println("log file:",logFileName)
    logFile,_ = os.Create(logFileName)
  }

  return logFile

}
