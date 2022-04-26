package main

const port = ":8080"

func main() {
	server := NewServer(port)
	server.Run()
}
