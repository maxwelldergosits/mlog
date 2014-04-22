package main

import "mlog"


func main() {


  levels := []string{"debug","network","general"} // These are the enabled levels

  log := mlog.Create(levels,"",true,false) // create logging system that disables the file and enables the levels in the slice above

  log.VPrintln("debug","This should Print")
  log.VPrintln("debug2","This should not Print")
  log.VPrintln("network","This should Print")
  log.NPrintln("This should Print")

  log.NPrintln("changing the logger!") // we have a new logger after this println
  levels = []string{}

  log = mlog.Create(levels,"",true,true) // create logging systems that prints all verbose and normal messages

  log.VPrintln("debug","This should Print")
  log.VPrintln("debug2","This should Print")
  log.VPrintln("network","This should Print")
  log.NPrintln("This should Print")

}
