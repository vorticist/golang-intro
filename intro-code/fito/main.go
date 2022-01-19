package main

import (
  "fmt"
  "os"
  "turkscheduler/services"
)

func main() {
  
  if len(os.Args) < 2 {
    panic(fmt.Sprintf("use %s [config_name]",os.Args[0])) 
  }
  //c,_:=config.Parse(fmt.Sprintf("config_json/%s_config.json",os.Args[1]))
  //fmt.Printf(c.ConnectionString())
  err:=services.Serve(fmt.Sprintf("config_json/%s_config.json",os.Args[1]))
  if err != nil{
    panic(err.Error())
  }
}
