//rest/register.do
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	UserName                   string `json:"userName,omitempty"`
	Password                   string `json:"password,omitempty"`
	Token                      string `json:"token,omitempty"`
	OrderNumber                string `json:"orderNumber"`
	Amount                     int    `json:"amount"`
	Currency                   int    `json:"currency,omitempty"`
	ReturnUrl                  string `json:"returnUrl"`
	FailUrl                    string `json:"failUrl,omitempty"`
	Description                string `json:"description,omitempty"`
	Ip                         string `json:"ip,omitempty"`
	Language                   string `json:"language,omitempty"`
	PageView                   string `json:"pageView,omitempty"`
	ClientId                   string `json:"clientId,omitempty"`
	MerchantLogin              string `json:"merchantLogin,omitempty"`
	Email                      string `json:"email,omitempty"`
	PostAddress                string `json:"postAddress,omitempty"`
	JsonParams                 string `json:"jsonParams,omitempty"`
	AdditionalOfdParams        string `json:"additionalOfdParams,omitempty"`
	SessionTimeoutSecs         int    `json:"sessionTimeoutSecs,omitempty"`
	ExpirationDate             string `json:"expirationDate,omitempty"`
	AutocompletionDate         string `json:"autocompletionDate,omitempty"`
	BindingId                  string `json:"bindingId,omitempty"`
	OrderBundle                string `json:"orderBundle,omitempty"`
	BillingPayerData           string `json:"billingPayerData,omitempty"`
	ShippingPayerData          string `json:"shippingPayerData,omitempty"`
	PreOrderPayerData          string `json:"preOrderPayerData,omitempty"`
	OrderPayerData             string `json:"orderPayerData,omitempty"`
	BillingAndShippingAddress  string `json:"billingAndShippingAddressMatchIndicator,omitempty"`
	Features                   string `json:"features,omitempty"`
	PrepaymentMdOrder          string `json:"prepaymentMdOrder,omitempty"`
	DynamicCallbackUrl         string `json:"dynamicCallbackUrl,omitempty"`
	FeeInput                   string `json:"feeInput,omitempty"`
	CardholderName             string `json:"cardholderName,omitempty"`
}

type RegisterResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	FormUrl      string `json:"formUrl,omitempty"`
	OrderId      string `json:"orderId,omitempty"`
}

func registerOrder(w http.ResponseWriter, r *http.Request) {
	// Check if it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Extract required fields
	orderNumber := r.FormValue("orderNumber")
	amountStr := r.FormValue("amount")
	returnUrl := r.FormValue("returnUrl")

	// Check required fields
	if orderNumber == "" || amountStr == "" || returnUrl == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Convert amount to integer
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Form the response (example)
	response := RegisterResponse{
		ErrorCode:    "0", // Assuming success
		FormUrl:      fmt.Sprintf("https://payment.gateway/%s", orderNumber),
		OrderId:      fmt.Sprintf("ORDER-%s", orderNumber),
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/rest/register.do", registerOrder)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//RegisterRequest — структура для приема данных запроса.
//RegisterResponse — структура ответа с полями ErrorCode, FormUrl и OrderId.
//Функция registerOrder:
//Проверяет, что это POST-запрос.
//Парсит данные формы.
//Проверяет наличие обязательных параметров: orderNumber, amount, returnUrl.
//Преобразует сумму заказа в целое число и проверяет на корректность.
//Возвращает JSON-ответ с formUrl и orderId.
