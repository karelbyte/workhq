package setup

import (
	"elpuertodigital/workhq/dbg"
	"fmt"
	"os"
)

func Migrate() {
	env := os.Getenv("ENV")
	if env == "local" {
		for _, entity := range dbg.SchemeEntities {
			_, err := dbg.DB().Exec(entity)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		fmt.Println("---------------------")
		fmt.Println("All migration success!")
	}

}
