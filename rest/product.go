package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Beelzebub0/redis-demo/lib"
)

func Product(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	switch r.Method {
	case "GET":
		redClient := lib.GetRedisConn()
		result, err := redClient.Get(ctx, "product").Result()
		if err != nil {
			http.Error(w, "Error", 404)
			return
		}
		w.Write([]byte(result))

	case "POST":
		redClient := lib.GetRedisConn()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		defer r.Body.Close()

		// SetEX --> Set key to hold the string value and set key to timeout after a given number of seconds
		// SetNX --> Set key to hold string value if key does not exist. In that case, it is equal to SET. When key already holds a value, no operation is performed. SETNX is short for "SET if Not eXists".
		_, err = redClient.SetEX(ctx, "product", body, time.Hour*24).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		fmt.Println("suzz")
		w.Write([]byte(body))
	}
}
