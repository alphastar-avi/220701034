package main

import(
	"fmt"
	"net/http"
	ser "urlshortner/utils"
)

func main(){
    ser.SetupRoute()
	fmt.Println("server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

