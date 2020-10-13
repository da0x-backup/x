package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func call(args []string) (string, error) {
	cmd := exec.Command("python3", "script.py", args[0])
	stdout, err := cmd.Output()
	if err != nil {
		println(err)
		return "", err
	}
	return string(stdout), nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	args, present := query["args"]
	if !present || len(args) <= 0 {
		log.Println("args not present", args)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Query parameter (args) is empty")
		return
	}
	output, err := call(args)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, output)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
