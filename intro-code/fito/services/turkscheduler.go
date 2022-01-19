package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"turkscheduler/config"
        "github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type InMemoryScheduledItem struct {
  ID string 
  WhenToLaunch time.Time
  Launched bool
}

type TurkScheduler struct {
  Connection *sql.DB
  NumberOfWorkers int 
  ScheduledItems []InMemoryScheduledItem
  itemFeedChannel chan InMemoryScheduledItem
}

func NewScheduler(config *config.WholeConfig) (scheduler *TurkScheduler ,err error){
  var ret TurkScheduler
  ret.Connection,err = sql.Open("postgres",config.ConnectionString())
  if err != nil {
    return
  }
  scheduler = &ret
  return
}

func (scheduler *TurkScheduler) PersistItem(id string ,description string, user_ids []string) error{

  _,err:=scheduler.Connection.Exec("INSERT INTO scheduled_items.scheduled_items VALUES($1,$2,$3)",
                            &id,&description,pq.Array(&user_ids))
  if err != nil{
    panic(err.Error())
  }
  return nil
}

func (scheduler *TurkScheduler) FindInvalidUser(user_ids []string)(invalid string, allok bool) {
  n:=0
  for _,v := range(user_ids){
    row:= scheduler.Connection.QueryRow("SELECT COUNT (*) FROM users.users WHERE user_id::text=$1",v)
    err:=row.Scan(&n)
    if err != nil{
      panic(err.Error())
    }
    if n ==0 {
      return v,false
    }
  }
  return "",true
}

func (scheduler *TurkScheduler ) TakeAction(item_id string) error{
  row:= scheduler.Connection.QueryRow("SELECT id,description,users::text FROM scheduled_items.scheduled_items")
  var id string 
  var users []string
  var description string
  err:=row.Scan(&id,&description,pq.Array(&users))
  if err != nil{
    return err
  }
  var mails []string  
  log.Printf("there are %d\n",len(users))
  rows,err := scheduler.Connection.Query("SELECT email FROM users.users WHERE user_id =any( $1)",pq.Array(&users))
  if err != nil{
    return err
  }
  mail:=""
  for rows.Next(){
    rows.Scan(&mail)
    mails=append(mails,mail)
  }
  outputUUID := uuid.New().String()
  _,err=scheduler.Connection.Exec("INSERT INTO output.output VALUES($1,$2,$3)",
                            &outputUUID,&description,pq.Array(&mails))

  return err
}

func (scheduler *TurkScheduler) Launch (pollingPeriod time.Duration, input chan InMemoryScheduledItem){
  scheduler.itemFeedChannel = input
   for {
     select{
     case v:=<-input:
       fmt.Printf("adding task")
       scheduler.ScheduledItems=append(scheduler.ScheduledItems,v)
     case <-time.After(time.Second*pollingPeriod):
       now :=time.Now()
       for i,v := range(scheduler.ScheduledItems){
         if (!v.Launched && now.After(v.WhenToLaunch)){
           log.Printf("I  have to launch item %s\n",v.ID)
           scheduler.ScheduledItems[i].Launched=true
           err:=scheduler.TakeAction(v.ID)
           if err != nil{
             fmt.Printf(err.Error())
           }
         }
       }
     }
  }
}
