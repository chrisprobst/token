package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

const (
	unassigned = -1
	chris      = 0
	andreas    = 1
	kalman     = 2
)

var header = `
<html>
<head>
</head>
<body>
Possible states:<br>
-----------------<br>
unassigned = -1<br>
chris = 0<br>
andreas = 1<br>
kalman = 2<br>
<br>
Actions:<br>
-----------------<br>
<a href="/">Refresh</a><br>
<a href="/release">Release Token</a><br>
<a href="/chris">Lock for Chris</a><br>
<a href="/andreas">Lock for Andreas</a><br>
<a href="/kalman">Lock for Kalman</a><br>
<br>
Actual state:<br>
-----------------<br>`

var footer = `
</body>
</html>
`

var token = unassigned
var m sync.Mutex

func root(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, header)
	fmt.Fprintf(w, "Token Owner: %d", token)
	fmt.Fprintln(w, footer)
}

func release(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	if token != -1 {
		token = -1
	}

	http.Redirect(w, r, "/", http.StatusFound);
}

func lockForChris(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	w.Header().Set("Content-Type", "text/html")

	if token != -1 && token != chris {
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d [Already locked, wait for release!]", token)
		fmt.Fprintln(w, footer)
	} else {
		token = chris
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d", token)
		fmt.Fprintln(w, footer)
	}
}

func lockForAndreas(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	w.Header().Set("Content-Type", "text/html")

	if token != -1 && token != andreas {
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d [Already locked, wait for release!]", token)
		fmt.Fprintln(w, footer)
	} else {
		token = andreas
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d", token)
		fmt.Fprintln(w, footer)
	}
}

func lockForKalman(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()
	w.Header().Set("Content-Type", "text/html")

	if token != -1 && token != kalman {
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d [Already locked, wait for release!]", token)
		fmt.Fprintln(w, footer)
	} else {
		token = kalman
		fmt.Fprintln(w, header)
		fmt.Fprintf(w, "Token Owner: %d", token)
		fmt.Fprintln(w, footer)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", root)
	http.HandleFunc("/release", release)
	http.HandleFunc("/chris", lockForChris)
	http.HandleFunc("/andreas", lockForAndreas)
	http.HandleFunc("/kalman", lockForKalman)
	http.ListenAndServe(":8669", nil)
}
