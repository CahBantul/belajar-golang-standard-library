package main

import "flag"

func main() {
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.String("port", "5432", "database port")

	flag.Parse()

	println("Connecting to database...")
	println("username", *username)
	println("password", *password)
	println("host", *host)
	println("port", *port)
}