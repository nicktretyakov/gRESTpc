package main

import (
	"encoding/json"
	//"fmt"
	//"log"
	"net/http"
	"net/url"
	"io/ioutil"
)

// Структура для запроса отмены заказа
type ReverseRequest struct {
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	OrderID       string `json:"orderId,omitempty"`
	OrderNumber   string `json:"orderNumber,omitempty"`
	Language      string `json:"language,omitempty"`
	JsonParams    string `json:"jsonParams,omitempty"`
	MerchantLogin string `json:"merchantLogin,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// Структура для ответа от API
type ReverseResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	UserMessage  string `json:"userMessage"`
}

func reverseOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	var req ReverseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Создание параметров для отправки запроса
	data := url.Values{}
	data.Set("userName", req.UserName)
	data.Set("password", req.Password)
	if req.OrderID != "" {
		data.Set("orderId", req.OrderID)
	}
	if req.OrderNumber != "" {
		data.Set("orderNumber", req.OrderNumber)
	}
	if req.Language != "" {
		data.Set("language", req.Language)
	}
	if req.JsonParams != "" {
		data.Set("jsonParams", req.JsonParams)
	}
	if req.MerchantLogin != "" {
		data.Set("merchantLogin", req.MerchantLogin)
	}
	if req.Amount != "" {
		data.Set("amount", req.Amount)
	}
	if req.Currency != "" {
		data.Set("currency", req.Currency)
	}

	// Отправка POST-запроса к внешнему API
	apiURL := "https://external-server.com/rest/reverse.do"
	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		http.Error(w, "Failed to make request to external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа от внешнего API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Преобразование ответа в JSON и отправка обратно клиенту
	var reverseResp ReverseResponse
	err = json.Unmarshal(body, &reverseResp)
	if err != nil {
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reverseResp)
}

//func main() {
//	http.HandleFunc("/rest/reverse.do", reverseOrder)
//
//	fmt.Println("Server is running on port 8080...")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//ReverseRequest содержит параметры для отправки запроса на отмену заказа.
//ReverseResponse описывает структуру ответа от внешнего API.
//Функция reverseOrder:
//Проверяет, что запрос отправлен методом POST.
//Принимает параметры в формате JSON, преобразует их в URL-закодированные параметры и отправляет на внешний сервер с помощью http.PostForm.
//Обрабатывает ответ и отправляет его обратно клиенту в виде JSON.
//Маршрут /rest/reverse.do:
//Принимает POST-запросы для отмены заказов.
//Сервер запускается на порту 8080 и обрабатывает запросы.
