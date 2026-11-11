# Defer
The defer keyword is a fairly unique fature of Go. It allows a function to be executed
automatically just before its enclosing function returns.

The deferred call's arguments are evaluated immediately, but the functions call is not executed until the surrounding function returns.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.

For example:

func GetUsername(dstName, srcName string) (username string, err error) {
	// Open a connection to a database
	conn, _ := db.Open(srcName)
	
	// Close the connection *anywhere* the GetUsername function returns
	
	defer conn.Close()
	
	username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}
	
	// The defer statement is auto=executed if we return here
	return username, nil
}


In the above example, the conn.Close() function is not called here:

defer conn.Close()


It's called:

// here
return "", err

// or here
return username, nil

#Multiple Defers

The location of a defer statement inside a function matters. The deferred call is registered at the point
where defer is executed, and it will run when the function returns. If you have multiple 
defer statements in a single function, they are executed in last-in, first-out order (the last deferred call runs first)

For example, you'd want to close a file before trying to remove it:

func CreateTempFile() {
	f, _ := os.Create("temp-42.txt")
	defer os.Remove(f.Name())
	defer f.Close()
	
	fmt.Fprintln(f, "How many roads must a man walk down?")
}
