package main

func main() {
	server := NewServer(":7070")
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Listen()
}
