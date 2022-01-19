package controllers

import (
	"net/http"
	"github.com/gorilla/sessions"
        "github.com/gorilla/mux"
)


type EndPoint struct {
  JSONBody bool
  Secure   bool 
  JSONResponse bool 
  fn func (http.ResponseWriter ,*http.Request)
  CookieStore *sessions.CookieStore
  router *mux.Router
  url string
}

func NewEndPoint(router *mux.Router, url string ) * EndPoint{
  var ret EndPoint 
  ret.router = router
  ret.url = url 
  return &ret 
}

func (endpoint *EndPoint ) UseJSONBody() *EndPoint{
  endpoint.JSONBody=true
  return endpoint
}

func (endpoint *EndPoint ) IsSecure() *EndPoint{
  endpoint.Secure=true
  return endpoint
}

func (endpoint *EndPoint ) UseJSONResponse() *EndPoint{
  endpoint.JSONResponse=true
  return endpoint
}
func (endpoint *EndPoint) Callback (fn func (http.ResponseWriter ,*http.Request)) *EndPoint{
  endpoint.fn =fn;
  return endpoint
}

func (endpoint *EndPoint ) Call( res http.ResponseWriter , req *http.Request){
 if endpoint.fn == nil{
   panic("unseted callback")
   return
 }
 if endpoint.JSONResponse{
   res.Header().Set("Content-Type", "application/json")
 }
 endpoint.fn(res, req)  
}

func (endpoint *EndPoint) Build(){
  endpoint.router.HandleFunc(endpoint.url,endpoint.fn);
}
