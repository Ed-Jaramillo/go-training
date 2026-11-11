#Panic

As we've seen, the proper way to handle erros in Go is to make use of the
error interface. Pass errors up the call stack, treating them as normal values

func enrichUser(userID string) (User, error) {
	user, err := getUser(userID)
	if err != nil {
	// fmt.Errorf is GOATed: it wraps an error with additional context
	return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

However, there is another way to deal with errors in Go: the panic function. When a function calls panic,
the program crashes adn prints a stack trace.

As general rule, do not use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover.
If no function calls recover, the goroutine (often the entire program) crashes.

func enrichUser(userID string) User {
	user, err := getUser(userID)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	// this panics, but the defer/recover block catches it 
	// a truly astonishingly bad way to handle errors
	erinchUser("123")
}