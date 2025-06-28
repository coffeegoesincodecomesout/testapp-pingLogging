package main

import (
   "fmt"
   "net/http"
   "log/slog"
   "os" 
)

func main() {
   http.HandleFunc("/ping", func (w http.ResponseWriter, req *http.Request){
  
   fmt.Fprintf(w,"pong")

   logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
   logger.Info("ping...pong...")

   })

   logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))          
   logger.Info("Server is starting... ")              

   http.ListenAndServe(":8090", nil)
}
