package controllers

import (
	"encoding/json"
	"net/http"

	"github.com.br/rodrygo262/hackaton_golang/models"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	provisioning := models.Provisioning{}
	err := json.NewDecoder(r.Body).Decode(&provisioning)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {

		if provisioning.OrderId == "" {
			http.Error(w, "orderId obrigatório", http.StatusBadRequest)
			return
		}

		if provisioning.ClientId == "" {
			http.Error(w, "clientId obrigatório", http.StatusBadRequest)
			return
		}

		if provisioning.SKU == "" {
			http.Error(w, "SKU obrigatório", http.StatusBadRequest)
			return
		}

		_, err := provisioning.Create_prov()

		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		} else {
			// mandar pro server de eng
		}

		http.Redirect(w, r, "/", http.StatusOK)
	} else {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	prov := models.Provisioning{}
	provisionings, err := prov.List_all_provisiongs()

	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	jData, err := json.Marshal(provisionings)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
