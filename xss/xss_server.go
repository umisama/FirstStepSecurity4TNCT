package main

import (
		"fmt"
		"net/http"
	)

var resp_template = `<html>
	<body>
		<p>あなたは %s と入力しましたね。</p>
	</body>
</html>`

var form_template = `<html>
	<body>
		<form action="hackme" method="post">
			<input type="text" name="value"></input>
			<input type="submit"></input>
			</form>
	</body>
</html>
`

func XssCapableHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("X-XSS-Protection", "0")
	fmt.Fprintf( w, resp_template, req.FormValue("value") )
}

func XssServeHandler( w http.ResponseWriter, req *http.Request ) {
	fmt.Fprintf( w, form_template )
}

func main() {
	http.HandleFunc("/hackme", XssCapableHandler )
	http.HandleFunc("/", XssServeHandler )
	http.ListenAndServe(":8080", nil)
}
