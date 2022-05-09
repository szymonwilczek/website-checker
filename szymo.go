package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input: %v\n", err)
	}

}

func checkDomain(domain string) {
	u, err := url.Parse(domain)

	scheme := u.Scheme                          // string -> "http" | "https"
	opaque := u.Opaque                          // string -> "scheme:opaque"
	user := u.User.String()                     // -> *Userinfo -> converted to String type -> "user:password" 	// -> string -> "password"
	host := u.Host                              // string -> "host:port"
	path := u.Path                              // string -> "/path"
	query := u.RawQuery                         // string -> "?query"
	fragment := u.Fragment                      // string -> "#fragment"
	forceQuery := u.ForceQuery                  //bool -> forced | not forced -> true/false
	hostname := u.Hostname()                    // string -> example: Parsed http://example.com -> Output will be: example.com
	wholeURL := u.String()                      // string -> "scheme:opaque?query#fragment"
	pathUnescape, _ := url.PathUnescape(domain) // string -> "/path"

	fmt.Println("Looking for:", domain)
	fmt.Println("Scheme:", scheme)
	fmt.Println("Opaque:", opaque)
	fmt.Println("User:", user)
	fmt.Println("Host:", host)
	fmt.Println("Path:", path)
	fmt.Println("Query:", query)
	fmt.Println("Fragment:", fragment)
	fmt.Println("ForceQuery:", forceQuery)
	fmt.Println("Hostname:", hostname)
	fmt.Println("WholeURL:", wholeURL)
	fmt.Println("Looked query:", pathUnescape)

	//Check availability of domain
	_, err = net.LookupHost(hostname)
	if err != nil {
		fmt.Println("Availability: ", "Domain is not available")
		return
	}
	fmt.Println("Availability: ", "Domain is available")

	//Check SSL certificate

	user2 := url.UserPassword(user, "password")

	if scheme == "https" {
		fmt.Println("SSL certificate: ", "Certificate is valid")
		fmt.Println("Password: ", user2)
		fmt.Println("Visible to website URL: ", u.Redacted(), "Your password kevlar is only available due to SSL certificate. Remember this.")

	} else {
		fmt.Println("SSL certificate: ", "Certificate is not valid! Watch out doing anything on this domain! Remember to cover your passwords!")
		fmt.Println("Password: ", user)
		fmt.Println("Visible to website URL: ", wholeURL, "Your password here can be leaked! Never give your password to non-SSL website.")
	}

}
