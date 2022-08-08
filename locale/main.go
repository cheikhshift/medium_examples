package main

import (
	"golang.org/x/text/message"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"

	"fmt"
)

func main(){


    message.SetString(
		language.French,
		"The %m car",
    	"La voiture %m",
    )

    message.SetString(
		language.French,
		"green",
    	"vert",
    )
	
    message.Set(language.English, "You are %d minute(s) late.",
    plural.Selectf(1, "",
        plural.One, "You are 1 minute late.",
        plural.Other, "You are %d minutes late."),
    )


	p := message.NewPrinter(language.French)
	fmt.Println( p.Sprintf("The %m car", "green") )
	fmt.Println( p.Sprintf("The %m car", "red")  )


	p2 := message.NewPrinter(language.English)
	fmt.Println( p2.Sprintf("You are %d minute(s) late.", 1 )  )
	fmt.Println( p2.Sprintf("You are %d minute(s) late.", 5 )  )

}