package main

type authenicationInfo struct {
	username string
	password string	
}


// create the method below

func (a authenticationInfo) getBasicAuth() string {
	return "Authorization: Baic" + " " + a.username + ":" + a.password 
}