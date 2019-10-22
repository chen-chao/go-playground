package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(loc)
	fmt.Println(now.In(loc))

	shanghai := time.FixedZone("Asia/Shanghai", +8*60*60)
	shanghaitime := now.In(shanghai)
	fmt.Println(shanghaitime)
}
