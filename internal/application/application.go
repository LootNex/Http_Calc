package application

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LootNex/Http_Calc/pkg/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Results struct {
	Result float64 `json:"result"`
}

type Errors struct {
	Error string `json:"error"`
}

type Config struct{
	Addr string
}

func ConfigFromEnv() *Config{
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == ""{
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application{
	return &Application{
		config: ConfigFromEnv(),
	}
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {

	request := new(Request)

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&Errors{Error: "Internal server error"})
		return
	}

	res, err := calculator.Calc(string(request.Expression))

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(&Errors{Error: "Expression is not valid"})
		return
	}

	result := Results{Result: res}
	err = json.NewEncoder(w).Encode(&result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&Errors{Error: "Internal server error"})
		return
	}
}

func (a Application) RunServer() error {

	http.HandleFunc("/api/v1/calculate", CalculatorHandler)

	return http.ListenAndServe(":"+a.config.Addr, nil)
}

func (a *Application) Run() error {
	for {

		log.Println("введите выражение")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("ошибка при чтении из консоли")
		}
		
		text = strings.TrimSpace(text)
		
		if text == "exit" {
			log.Println("выполнено успешно")
			return nil
		}
		
		result, err := calculator.Calc(text)
		if err != nil {
			log.Println(text, "ошибка:", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}
