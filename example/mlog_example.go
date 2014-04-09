package main

import "mlog"


func main() {


  levels := []string{"debug","network","general"}

  log := mlog.Create(levels,"",true,false)

  log.VPrintln("debug","This should Print")
  log.VPrintln("debug2","This should not Print")
  log.VPrintln("network","This should Print")
  log.NPrintln("This should Print")

  log.NPrintln("NEW TEST")
  levels = []string{}

  log = mlog.Create(levels,"",true,true)

  log.VPrintln("debug","This should Print")
  log.VPrintln("debug2","This should Print")
  log.VPrintln("network","This should Print")
  log.NPrintln("This should Print")
}
