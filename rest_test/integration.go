package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Константы для подключения к платёжному шлюзу
const (
	username   = "USERNAME"
	password   = "PASSWORD"
	gatewayURL = "https://server/payment/rest/"
	returnURL  = "http://your.site/rest.php"
)

func main() {
	http.HandleFunc("/rest", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Функция взаимодействия с платёжным шлюзом
func gateway(method string, data url.Values) (map[string]interface{}, error) {
	resp, err := http.PostForm(gatewayURL+method, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Обработчик формы
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Query().Get("orderId") == "" {
		// Вывод формы на экран
		fmt.Fprintf(w, `
			<form method="post" action="/rest">
				<label>Order number</label><br />
				<input type="text" name="orderNumber" /><br />
				<label>Amount</label><br />
				<input type="text" name="amount" /><br />
				<button type="submit">Submit</button>
			</form>
		`)
	} else if r.Method == http.MethodPost {
		// Обработка данных из формы
		orderNumber := r.FormValue("orderNumber")
		amount := r.FormValue("amount")

		data := url.Values{
			"userName":    {username},
			"password":    {password},
			"orderNumber": {url.QueryEscape(orderNumber)},
			"amount":      {url.QueryEscape(amount)},
			"returnUrl":   {returnURL},
		}

		response, err := gateway("register.do", data)
		if err != nil {
			http.Error(w, "Ошибка взаимодействия с шлюзом", http.StatusInternalServerError)
			return
		}

		if errorCode, ok := response["errorCode"]; ok {
			// В случае ошибки вывести её
			fmt.Fprintf(w, "Ошибка #%v: %v", errorCode, response["errorMessage"])
		} else {
			// В случае успеха перенаправить на платёжную форму
			http.Redirect(w, r, response["formUrl"].(string), http.StatusSeeOther)
		}
	} else if r.Method == http.MethodGet && r.URL.Query().Get("orderId") != "" {
		// Обработка данных после формы
		orderId := r.URL.Query().Get("orderId")

		data := url.Values{
			"userName": {username},
			"password": {password},
			"orderId":  {orderId},
		}

		response, err := gateway("getOrderStatus.do", data)
		if err != nil {
			http.Error(w, "Ошибка взаимодействия с шлюзом", http.StatusInternalServerError)
			return
		}

		// Вывод кода ошибки и статуса заказа
		fmt.Fprintf(w, `
			<b>Error code:</b> %v<br />
			<b>Order status:</b> %v<br />
		`, response["ErrorCode"], response["OrderStatus"])
	}
}
//Используется http.PostForm для отправки POST-запросов с параметрами.
//Данные передаются в формате url.Values, который Go автоматически конвертирует в query-параметры для POST запроса.
//Ответы обрабатываются как JSON с помощью json.Unmarshal.

