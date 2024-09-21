package main

/*


// lets solve the sleeping barber problem using go routines and channels
// the problem is as follows: there is a barber shop with one barber and n chairs for waiting customers so we can have n+1 customers at most in the shop

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping // when you get something from ping put inot s

		pong <- fmt.Sprintf("%s", strings.ToUpper(s)) // put s into pong
	}
}

// ping takes the varibale and pong prints the variable in uppercase so we can see the output in uppercase]
// clearly we can see that the ping and pong are two channels that are used to communicate between the two go routines
// the ping channel is used to send the data to the shout function and the pong channel is used to get the data from the shout function
// the shout function is a go routine that is running in the background and it is listening to the ping channel and when it gets the data it puts the data into the pong channel
// the main function is the main go routine that is running in the foreground and it is listening to the pong channel and when it gets the data it prints the data in the console
// the main function is also sending the data to the ping channel so that the shout function can get the data and put it into the pong channel
// so the main function is sending the data to the shout function and the shout function is sending the data to the main function

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	//time.Sleep(10 * time.Second)
	fmt.Println("type something to enter and press q to quit")

	for {
		fmt.Print("-->")

		var input string
		_, _ = fmt.Scanln(&input)
		if input == strings.ToLower("q") {
			break
		}

		ping <- input

		response := <-pong
		fmt.Println("response: ", response)

	}

	fmt.Println("done all channels are closed now ")
}
*/
