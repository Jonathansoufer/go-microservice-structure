package main

import (
	"fmt"
	"log"
)

func main()	{
	svc := NewLoggingService(NewGetCatFactService("https://cat-fact.herokuapp.com/facts/random"))

	apiserver := NewApiserver(svc)
	if e := apiserver.Start(":8080"); e != nil {
		log.Fatal(e)
	}

	fmt.Println("done")

}