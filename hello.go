funcc main() {
  http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, 세계!")
  })
  log.Fatall(http.ListenAndServer(":8080",nil))
}
