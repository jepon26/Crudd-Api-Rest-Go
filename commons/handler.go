package commons

import "net/http"

func SendResponse(write http.ResponseWriter, status int, data[]byte){
	write.Header().Set("Content-Type", "applicatiom/json")
	write.WriteHeader(status);
	write.Write(data);
}
