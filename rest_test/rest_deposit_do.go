package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DepositRequest struct {
	UserName                        string `json:"userName"`
	Password                        string `json:"password"`
	OrderId                         string `json:"orderId"`
	Amount                          int    `json:"amount"`
	Language                        string `json:"language,omitempty"`
	JsonParams                      string `json:"jsonParams,omitempty"`
	DepositItems                    string `json:"depositItems,omitempty"`
	Agent                           string `json:"agent,omitempty"`
	SupplierPhones                  string `json:"supplierPhones,omitempty"`
	DepositType                     int    `json:"depositType,omitempty"`
	Currency                        string `json:"currency,omitempty"`
	MultipleCompletionOrderDescription string `json:"multipleCompletionOrderDescription,omitempty"`
}

type DepositResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	UserMessage  string `json:"userMessage"`
}

func depositOrder(apiUrl string, requestData DepositRequest) (*DepositResponse, error) {
	// Конвертируем структуру в JSON
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

	// Логируем ответ для отладки
	log.Printf("Response status: %s", resp.Status)
	log.Printf("Response body: %s", body)

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Парсим JSON-ответ
	var response DepositResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &response, nil
}

func main() {
	apiUrl := "https://example.com/rest/deposit.do"

	// Данные для запроса
	requestData := DepositRequest{
		UserName:   "merchant_login",
		Password:   "merchant_password",
		OrderId:    "123456789012345678901234567890123456", // Пример Order ID
		Amount:     10000,                                 // Сумма в валюте заказа (например, 100 рублей = 10000 копеек)
		Language:   "ru",
		JsonParams: "{\"Имя1\": \"Значение1\", \"Имя2\": \"Значение2\"}",
		DepositItems: `[{"positionId":1,"name":"Товар","quantity":{"value":1,"measure":"шт."},"itemCode":"12345","itemAmount":10000,"itemPrice":10000,"tax":{"taxType":0}}]`,
		Agent:      "{\"Имя1\": \"Значение1\", \"Имя2\": \"Значение2\"}",
		SupplierPhones: "+79161234567",
		DepositType: 0,
		Currency:    "643", // Код валюты (643 = рубли)
	}

	// Отправляем запрос
	response, err := depositOrder(apiUrl, requestData)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Выводим данные ответа
	fmt.Printf("Order deposit response: %+v\n", response)
}
//DepositRequest для параметров запроса, включая обязательные (UserName, Password, OrderId, Amount) и дополнительные параметры, такие как Language, JsonParams, //DepositItems и другие.
//Преобразует структуру запроса в JSON.
//Создает HTTP-запрос методом POST и отправляет его на указанный URL.
//Обрабатывает ответ, проверяет статус и выводит результат.
//Логирует запрос и ответ для отладки.
