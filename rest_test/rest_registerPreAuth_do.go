package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Константы для подключения к платёжному шлюзу
const (
	gatewayURL = "https://server/payment/rest/"
	username   = "USERNAME"
	password   = "PASSWORD"
)

// Структура для параметров запроса
type PreAuthRequest struct {
	UserName              string `json:"userName"`
	Password              string `json:"password"`
	Token                 string `json:"token,omitempty"`
	OrderNumber           string `json:"orderNumber"`
	Amount                int    `json:"amount"`
	Currency              int    `json:"currency,omitempty"`
	ReturnUrl             string `json:"returnUrl"`
	FailUrl               string `json:"failUrl,omitempty"`
	Description           string `json:"description,omitempty"`
	IP                    string `json:"ip,omitempty"`
	Language              string `json:"language,omitempty"`
	PageView              string `json:"pageView,omitempty"`
	ClientID              string `json:"clientId,omitempty"`
	MerchantLogin         string `json:"merchantLogin,omitempty"`
	Email                 string `json:"email,omitempty"`
	PostAddress           string `json:"postAddress,omitempty"`
	JsonParams            string `json:"jsonParams,omitempty"`
	AdditionalOfdParams   string `json:"additionalOfdParams,omitempty"`
	SessionTimeoutSecs    int    `json:"sessionTimeoutSecs,omitempty"`
	ExpirationDate        string `json:"expirationDate,omitempty"`
	AutocompletionDate    string `json:"autocompletionDate,omitempty"`
	BindingID             string `json:"bindingId,omitempty"`
	OrderBundle           string `json:"orderBundle,omitempty"`
	BillingPayerData      string `json:"billingPayerData,omitempty"`
	ShippingPayerData     string `json:"shippingPayerData,omitempty"`
	PreOrderPayerData     string `json:"preOrderPayerData,omitempty"`
	OrderPayerData        string `json:"orderPayerData,omitempty"`
	BillingShippingMatch  string `json:"billingAndShippingAddressMatchIndicator,omitempty"`
	Features              string `json:"features,omitempty"`
	PrepaymentMdOrder     string `json:"prepaymentMdOrder,omitempty"`
	DynamicCallbackUrl    string `json:"dynamicCallbackUrl,omitempty"`
	FeeInput              string `json:"feeInput,omitempty"`
	CardholderName        string `json:"cardholderName,omitempty"`
}

// Структура для ответа от платежного шлюза
type GatewayResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	FormURL      string `json:"formUrl"`
	OrderID      string `json:"orderId"`
}

// Функция для взаимодействия с платёжным шлюзом
func gatewayRequest(method string, data url.Values) (GatewayResponse, error) {
	resp, err := http.PostForm(gatewayURL+method, data)
	if err != nil {
		return GatewayResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GatewayResponse{}, err
	}

	var result GatewayResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return GatewayResponse{}, err
	}

	return result, nil
}

// Обработчик POST запроса на предавторизацию
func registerPreAuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	var preAuthRequest PreAuthRequest
	err := json.NewDecoder(r.Body).Decode(&preAuthRequest)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}

	// Проверка обязательных параметров
	if preAuthRequest.OrderNumber == "" || preAuthRequest.Amount == 0 || preAuthRequest.ReturnUrl == "" {
		http.Error(w, "Параметры orderNumber, amount и returnUrl обязательны", http.StatusBadRequest)
		return
	}

	// Подготовка данных для отправки на платёжный шлюз
	data := url.Values{}
	data.Set("userName", username)
	data.Set("password", password)
	data.Set("orderNumber", preAuthRequest.OrderNumber)
	data.Set("amount", strconv.Itoa(preAuthRequest.Amount))
	data.Set("returnUrl", preAuthRequest.ReturnUrl)

	// Опциональные параметры
	if preAuthRequest.Currency != 0 {
		data.Set("currency", strconv.Itoa(preAuthRequest.Currency))
	}
	if preAuthRequest.FailUrl != "" {
		data.Set("failUrl", preAuthRequest.FailUrl)
	}
	if preAuthRequest.Description != "" {
		data.Set("description", preAuthRequest.Description)
	}
	if preAuthRequest.IP != "" {
		data.Set("ip", preAuthRequest.IP)
	}
	if preAuthRequest.Language != "" {
		data.Set("language", preAuthRequest.Language)
	}

	// Отправка запроса на платёжный шлюз
	response, err := gatewayRequest("registerPreAuth.do", data)
	if err != nil {
		http.Error(w, "Ошибка при отправке запроса на платёжный шлюз", http.StatusInternalServerError)
		return
	}

	// Отправка ответа клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//func main() {
//	http.HandleFunc("/rest/registerPreAuth.do", registerPreAuthHandler)
//	fmt.Println("Сервер запущен на порту 8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//PreAuthRequest: Структура, представляющая параметры, которые будут переданы в запрос.
//GatewayResponse: Структура для обработки ответа от платёжного шлюза.
//gatewayRequest: Функция, которая выполняет POST запрос на платежный шлюз с использованием данных запроса.
//registerPreAuthHandler: Обработчик для POST запросов на /rest/registerPreAuth.do. Проверяет наличие обязательных параметров и отправляет запрос на шлюз.
//main: Функция запускает HTTP сервер на порту 8080 и регистрирует обработчик для маршрута /rest/registerPreAuth.do.
// curl -X POST http://localhost:8080/rest/registerPreAuth.do \
// -H "Content-Type: application/json" \
// -d '{
//    "orderNumber": "12345",
//    "amount": 5000,
//    "returnUrl": "https://your.site/success",
//    "failUrl": "https://your.site/fail"
//}'
