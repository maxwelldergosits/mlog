package mlog


import (
  "log"
  "os"
  "io"
)


type MLog struct {
  logLevels map[string]bool
  normalWriter io.Writer
  verboseWriter io.Writer
  all bool
}


func Create(levels []string, file_prefix string,disableLog bool, all bool) MLog {

  t := MLog{
    make(map[string]bool),
    nil,
    nil,
    all}
  for _,level := range levels {
    t.logLevels[level]=true
  }

  logFile := getFileWriter(file_prefix,disableLog)

  t.normalWriter = io.MultiWriter(logFile,os.Stdout)
  t.verboseWriter = io.MultiWriter(logFile)

  return t
}

func (t * MLog) VPrintln(priority string, a...interface{}){

  if t.logLevels[priority] || t.all{
    log.SetOutput(t.normalWriter)
  } else {
    log.SetOutput(t.verboseWriter)
  }

  b := []interface{}{priority+": "}
  c := append(b,a...)
  log.Println(c...)

}

func (t* MLog) VPrintf(priority string, format string, a...interface{}){

  if t.logLevels[priority]  || t.all {
    log.SetOutput(t.normalWriter)
  } else {
    log.SetOutput(t.verboseWriter)
  }


  log.Printf(priority+": "+format,a...)

}

func (t* MLog) NPrintf(format string, a...interface{}) {
    log.SetOutput(t.normalWriter)
    log.Printf(format,a...)
}
func (t* MLog) NPrintln(a...interface{}){
  log.SetOutput(t.normalWriter)
  log.Println(a...)
}
