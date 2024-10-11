package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Структура для запроса
type OrderStatusRequest struct {
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	OrderId       string `json:"orderId,omitempty"`
	OrderNumber   string `json:"orderNumber,omitempty"`
	Token         string `json:"token,omitempty"`
	Language      string `json:"language,omitempty"`
	MerchantLogin string `json:"merchantLogin,omitempty"`
}

// Структура для ответа
type OrderStatusResponse struct {
	ActionCode              int    `json:"actionCode"`
	ActionCodeDescription   string `json:"actionCodeDescription"`
	Amount                  int64  `json:"amount"`
	ErrorCode               string `json:"errorCode"`
	ErrorMessage            string `json:"errorMessage"`
	UserMessage             string `json:"userMessage"`
	OrderStatus             int    `json:"orderStatus"`
	OrderDescription        string `json:"orderDescription"`
	OrderNumber             string `json:"orderNumber"`
	PaymentAmountInfo       struct {
		ApprovedAmount int64  `json:"approvedAmount"`
		PaymentState   string `json:"paymentState"`
		TotalAmount    int64  `json:"totalAmount"`
	} `json:"paymentAmountInfo"`
}

// Функция для отправки запроса на получение статуса заказа
func getOrderStatus(apiUrl string, requestData OrderStatusRequest) (*OrderStatusResponse, error) {
	// Преобразуем запрос в JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request data: %v", err)
	}

	// Создаем POST-запрос
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP response: %v", err)
	}

	// Логируем статус ответа
	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", body)

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Парсим JSON-ответ
	var response OrderStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &response, nil
}

func main() {
	apiUrl := "https://example.com/rest/getOrderStatusExtended.do"

	// Данные для запроса
	requestData := OrderStatusRequest{
		UserName:    "merchant_login",
		Password:    "merchant_password",
		OrderId:     "123456789012345678901234567890123456", // Или используйте OrderNumber
		Language:    "ru",
		MerchantLogin: "merchant_login",
	}

	// Отправляем запрос
	response, err := getOrderStatus(apiUrl, requestData)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Выводим данные ответа
	fmt.Printf("Order status response: %+v\n", response)
}
//OrderStatusRequest — структура для запроса, которая включает параметры UserName, Password, а также опциональные параметры OrderId, OrderNumber, Token, Language, и MerchantLogin.
//OrderStatusResponse — структура для ответа, которая содержит поля, соответствующие ответу API, включая статус заказа, код действия и сумму.
//getOrderStatus — функция для отправки запроса с использованием метода POST, конвертации данных в JSON, отправки и обработки ответа.

