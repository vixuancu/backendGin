package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run the application.
	http.HandleFunc("/demo", demoFunc)
	log.Println("hello world")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
func demoFunc(w http.ResponseWriter, req *http.Request) {
	log.Printf("%+v", req)
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respone := map[string]string{
		"message": "Lap trinh golang",
		"info":    "vixuancu",
	}
	w.Header().Set("Content-Type", "application/json") //set giá trị trả về là "application/json"
	w.Header().Set("X-Courese", "Lap trinh Golang")    //set giá trị của header X-Courese
	//data, err := json.Marshal(respone)
	//if err != nil {
	//	http.Error(w, "Internal server error lỗi", http.StatusInternalServerError)
	//	return
	//}
	//w.Write(data)
	json.NewEncoder(w).Encode(respone) // Sử dụng json.NewEncoder để tự động mã hóa và ghi vào ResponseWriter
}
