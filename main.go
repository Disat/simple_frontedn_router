package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	//mux := http.NewServeMux()

	//// Specify the directory to serve files from
	//fs := http.Dir("/")
	//
	//// Create a custom file server handler
	//fileServer := &customFileServer{root: fs}
	//http.Handle("/", fileServer)
	//
	////filehandle := http.FileServer(http.Dir("./frontend"))
	////mux.Handle("/", NotFoundMiddleware(filehandle))

	//fileHandler := http.FileServer(http.Dir("./frontend"))

	// Create a custom handler to intercept 404 errors
	//customHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	// Serve the file using the file handler
	//	fileHandler.ServeHTTP(w, r)
	//	// Check if the status code is 404
	//	if _, ok := w.(http.Hijacker); !ok {
	//		log.Println("Webserver doesn't support hijacking")
	//		return
	//	}
	//	//w.WriteHeader(http.StatusNotFound)
	//	// Write your custom 404 error message
	//	w.Write([]byte("Custom 404 - Page Not Found"))
	//})
	customHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested file exists
		_, err := os.Stat(filepath.Join("./frontend", r.URL.Path))
		if err != nil {
			// File not found, respond with custom 404 message
			//http.Error(w, "Custom 404 - Page Not Found", http.StatusNotFound)
			http.ServeFile(w, r, "./frontend/index.html")
			return
		}

		// File exists, let the file server handle the request
		http.FileServer(http.Dir("./frontend")).ServeHTTP(w, r)
	})
	mux := http.NewServeMux()

	// Handle requests to the root URL with the custom handler
	mux.Handle("/", customHandler)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8181",
	}
	http.NotFoundHandler()

	server.ListenAndServe()
}

//type customFileServer struct {
//	root http.Dir
//}
//
//func (c *customFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// Convert request path to a cleaned version
//	cleanPath := filepath.Clean(r.URL.Path)
//
//	// Construct the full file path
//	fullPath := filepath.Join(string(c.root), cleanPath)
//
//	// Check if the requested file exists
//	_, err := os.Stat(fullPath)
//	if os.IsNotExist(err) {
//		// If file does not exist, handle 404 response
//		fmt.Fprintf(w, "adfasdfasdfas")
//		return
//	}
//
//	// Otherwise, serve the file using the standard file server
//	http.FileServer(c.root).ServeHTTP(w, r)
//}

//type cacheResponse struct {
//	cachestatus int
//	http.ResponseWriter
//}

//func (c *cacheResponse) WriteHeader(status int) {
//	c.cachestatus = status
//	c.ResponseWriter.WriteHeader(status)
//	if c.cachestatus == http.StatusNotFound {
//		c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
//	}
//}

//func (c *cacheResponse) Write(data []byte) (int, error) {
//
//	file, err := os.ReadFile("./frontend/index.html")
//	if err != nil {
//		return 0, err
//	}
//
//	// You can add your logic here to handle writing response body
//	// For example, you can clear the existing buffer and write your content
//	//c.ResponseWriter.WriteHeader(c.cachestatus) // Ensure status code is written
//	return c.ResponseWriter.Write(file)
//}

//func NotFoundMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Create a custom ResponseWriter to capture the status code
//		captureWriter := &cacheResponse{ResponseWriter: w}
//
//		// Call the next handler
//		next.ServeHTTP(captureWriter, r)
//
//		//Check if status code is 404
//		if captureWriter.cachestatus == http.StatusNotFound {
//			// Handle 404 error here
//			fmt.Println("404 Error: Page not found")
//
//		}
//	})
//}
