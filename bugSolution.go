func handleRequest(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            log.Printf("Recovered from panic: %v", r)
        }
    }()
    // ... some code that may panic ...
} 