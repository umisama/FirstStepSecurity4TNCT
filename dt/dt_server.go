package main

import (
		"io"
		"os"
		"fmt"
		"net/http"
	)

var form_template = `<html>
	<body>
		<form action="hackme" method="post">
			<input type="text" name="value"></input>
			<input type="submit"></input>
			</form>
	</body>
</html>
`

func DTravasalCapableHandler(w http.ResponseWriter, req *http.Request) {
	f, _ := os.Open( "www/" + req.FormValue("value"))
	io.Copy( w, f )
}

func XssServeHandler( w http.ResponseWriter, req *http.Request ) {
	fmt.Fprintf( w, form_template )
}

func main() {
	http.HandleFunc("/hackme", DTravasalCapableHandler )
	http.HandleFunc("/", XssServeHandler )
	http.ListenAndServe(":8080", nil)
}
