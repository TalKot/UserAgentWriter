package main

import (
	"fmt"
)

func main() {
	printLogo()
	settingConsumerSettings()
	startConsumer()
}

func printLogo() {
	fmt.Println("")
	fmt.Println("______      __   __ __      __ ")
	fmt.Println("/_  __/___ _/ /  / //_/___  / /_")
	fmt.Println("/ / / __ `/ /  / ,< / __ \/ __/")
	fmt.Println("/ / / /_/ / /  / /| / /_/ / /_  ")
	fmt.Println("/_/  \__,_/_/  /_/ |_\____/\__/  ")
	fmt.Println("")

}
