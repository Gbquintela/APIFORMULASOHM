package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct para armazenar os dados conforme as grandezas elétricas
type LeisDeOhm struct {
	Resistencia         float64 `json:"resistencia"`          // Resistência (ohms)
	Corrente            float64 `json:"corrente"`             // Corrente (amperes)
	Tensao              float64 `json:"tensao"`               // Tensão (volts)
	PotenciaVI          float64 `json:"potencia_vi"`          // Potência usando Tensão e Corrente (watts)
	Tempo               float64 `json:"tempo"`                // Tempo (segundos)
	EnergiaPT           float64 `json:"energia_pt"`           // Energia elétrica (joules) usando Potência e Tempo
	Resistividade       float64 `json:"resistividade"`        // Resistividade (ohm metros)
	Comprimento         float64 `json:"comprimento"`          // Comprimento (metros)
	Area                float64 `json:"area"`                 // Área da seção transversal (metros quadrados)
	ResistenciaEletrica float64 `json:"resistencia_eletrica"` // Resistência elétrica (ohms)
	PotenciaIR          float64 `json:"potencia_ir"`          // Potência usando Corrente e Resistência (watts)
	EnergiaVIT          float64 `json:"energia_vit"`          // Energia usando Tensão, Corrente e Tempo
	EnergiaCRT          float64 `json:"energia_crt"`          // Energia usando Corrente, Resistência e Tempo
	EnergiaVRT          float64 `json:"energia_vrt"`          // Energia usando Tensão, Resistência e Tempo
	Potencia            float64 `json:"potencia"`             // Potência total (watts)
	PotenciaVR          float64 `json:"potencia_vr"`          // Potência usando Corrente ao quadrado e Resistência (watts)
}

var dados LeisDeOhm

// CalcularTensao calcula a Tensão usando a Lei de Ohm (V = I * R)
func CalcularTensao(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if dados.Resistencia == 0 {
		http.Error(w, "Resistência não pode ser zero para calcular a tensão", http.StatusBadRequest)
		return
	}
	dados.Tensao = dados.Corrente * dados.Resistencia

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A tensão é %.2f volts", dados.Tensao),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularCorrente calcula a Corrente usando a Lei de Ohm (I = V / R)
func CalcularCorrente(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if dados.Resistencia == 0 {
		http.Error(w, "Resistência não pode ser zero para calcular a corrente", http.StatusBadRequest)
		return
	}
	dados.Corrente = dados.Tensao / dados.Resistencia

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A corrente é %.2f amperes", dados.Corrente),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularResistencia calcula a Resistência usando a Lei de Ohm (R = V / I)
func CalcularResistencia(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if dados.Corrente == 0 {
		http.Error(w, "Corrente não pode ser zero para calcular a resistência", http.StatusBadRequest)
		return
	}
	dados.Resistencia = dados.Tensao / dados.Corrente

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A resistência é %.2f ohms", dados.Resistencia),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularResistividade calcula a Resistividade (ρ = (R * A) / L)
func CalcularResistividade(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if dados.Resistividade <= 0 || dados.Comprimento <= 0 || dados.Area <= 0 {
		http.Error(w, "Os valores de resistividade, comprimento e área devem ser maiores que zero.", http.StatusBadRequest)
		return
	}
	dados.ResistenciaEletrica = dados.Resistividade * (dados.Comprimento / dados.Area)

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A resistência elétrica é %.2f", dados.ResistenciaEletrica),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularEnergia calcula a Energia Elétrica (E = P * t)
func CalcularEnergia(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.EnergiaPT = dados.Potencia * dados.Tempo

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A energia elétrica é %.2f", dados.EnergiaPT),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularEnergiaVIT calcula a Energia usando Tensão, Corrente e Tempo (E = V * I * t)
func CalcularEnergiaVIT(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.EnergiaVIT = dados.Tensao * dados.Corrente * dados.Tempo

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A energia usando tensão, corrente e tempo é %.2f", dados.EnergiaVIT),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularEnergiaCRT calcula a Energia usando Corrente, Resistência e Tempo (E = I^2 * R * t)
func CalcularEnergiaCRT(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.EnergiaCRT = (dados.Corrente * dados.Corrente) * dados.Resistencia * dados.Tempo

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A energia usando corrente, resistência e tempo é %.2f", dados.EnergiaCRT),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularEnergiaVRT calcula a Energia usando Tensão, Resistência e Tempo (E = (V^2) / R * t)
func CalcularEnergiaVRT(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.EnergiaVRT = (dados.Tensao * dados.Tensao * dados.Tempo) / dados.Resistencia

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A energia usando tensão, resistência e tempo é %.2f", dados.EnergiaVRT),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularPotenciaVI calcula a Potência usando Tensão e Corrente (P = V * I)
func CalcularPotenciaVI(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.PotenciaVI = dados.Tensao * dados.Corrente

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A potência usando tensão e corrente é %.2f watts", dados.PotenciaVI),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularPotenciaIR calcula a Potência usando Corrente e Resistência (P = I^2 * R)
func CalcularPotenciaIR(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.PotenciaIR = dados.Corrente * dados.Corrente * dados.Resistencia

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A potência usando corrente e resistência é %.2f watts", dados.PotenciaIR),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CalcularPotenciaVR calcula a Potência usando Corrente ao quadrado e Resistência (P = (I^2) * R)
func CalcularPotenciaVR(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dados.PotenciaVR = (dados.Corrente * dados.Corrente) * dados.Resistencia

	response := struct {
		Resposta string `json:"resposta"`
	}{
		Resposta: fmt.Sprintf("A potência usando corrente ao quadrado e resistência é %.2f watts", dados.PotenciaVR),
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Página inicial da API
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Leis de Ohm.\nUse os endpoints POST /CalcularTensao, /CalcularCorrente, /CalcularResistencia, /CalcularResistividade, /CalcularEnergia, /CalcularPotenciaVI, /CalcularPotenciaIR, /CalcularPotenciaVR para realizar cálculos elétricos.")
}

func main() {
	r := mux.NewRouter()

	log.Println("Iniciando API das Leis de Ohm")

	// Definindo os handlers para os endpoints
	r.HandleFunc("/homepage", homepage).Methods("GET")
	r.HandleFunc("/CalcularTensao", CalcularTensao).Methods("POST")
	r.HandleFunc("/CalcularCorrente", CalcularCorrente).Methods("POST")
	r.HandleFunc("/CalcularResistencia", CalcularResistencia).Methods("POST")
	r.HandleFunc("/CalcularResistividade", CalcularResistividade).Methods("POST")
	r.HandleFunc("/CalcularEnergia", CalcularEnergia).Methods("POST")
	r.HandleFunc("/CalcularEnergiaVIT", CalcularEnergiaVIT).Methods("POST")
	r.HandleFunc("/CalcularEnergiaCRT", CalcularEnergiaCRT).Methods("POST")
	r.HandleFunc("/CalcularEnergiaVRT", CalcularEnergiaVRT).Methods("POST")
	r.HandleFunc("/CalcularPotenciaVI", CalcularPotenciaVI).Methods("POST")
	r.HandleFunc("/CalcularPotenciaIR", CalcularPotenciaIR).Methods("POST")
	r.HandleFunc("/CalcularPotenciaVR", CalcularPotenciaVR).Methods("POST")

	// Iniciando o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
