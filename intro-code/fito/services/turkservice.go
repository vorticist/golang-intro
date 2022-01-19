package services 

import(
   "net/http"
   "turkscheduler/config"
   "turkscheduler/controllers"
   "github.com/gorilla/mux"  
   "time"
   "log"
   "encoding/json"
   "github.com/google/uuid"
)

type ScheduleRequest struct {
  Users []string
  WhenToLaunch string 
  Description string
}

func (scheduler *TurkScheduler) registerItem( res http.ResponseWriter , req *http.Request){

  var data ScheduleRequest
  err := json.NewDecoder(req.Body).Decode(&data)
  if err != nil {
    log.Print(err.Error())
    res.Write([]byte(err.Error()))
    return 
  }
  if len(data.Users) == 0 {
    res.Write([]byte("supply at least one user id"))
    return 
  }
  inv,allok := scheduler.FindInvalidUser(data.Users)
  if (!allok){
    res.Write([]byte("invalid UUID "+inv))
    return
  }

  structTime , err := time.Parse(time.RFC3339,data.WhenToLaunch)
  if err != nil{
    res.Write([]byte("Invalid date format use yyy-mm-ddTHH-MM:ss+(Timezone)"))
    return 
  }
  if data.Description == ""{
    res.Write([]byte("Description must contain an string"))
    return 
  }
  taskUUID := uuid.New().String()
  scheduler.PersistItem(taskUUID,data.Description,data.Users)
  scheduler.itemFeedChannel <- InMemoryScheduledItem{taskUUID,structTime,false}
  log.Printf("%s\n",structTime.String())
  res.Write([]byte(taskUUID))
}

func Serve(configPath string)(err error){
  
  conf,err := config.Parse(configPath)
  if err != nil {
    return 
  }

  scheduler,err :=NewScheduler(conf)
  if err != nil{
    return
  }
  
  exchangeChannel :=make(chan InMemoryScheduledItem)
  go func(){
    scheduler.Launch(1,exchangeChannel)
  }()
  log.Printf("The scheduler should be running"); 
  
  router := mux.NewRouter()
  controllers.NewEndPoint(router,"/schedule/").Callback(scheduler.registerItem).Build()
  srv := &http.Server{
		Handler: router,
		Addr:    conf.Serverbase,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
  log.Printf("ready to rock")
  err=srv.ListenAndServe()
  return 
}

