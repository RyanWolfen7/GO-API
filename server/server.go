package server

// Init runs router
func Init() {
	router := NewRouter()
	router.Run(":8000")
}
