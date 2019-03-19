package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kevinhohl/petstore/pkg/database"
	"github.com/kevinhohl/petstore/pkg/model"

	"github.com/julienschmidt/httprouter"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func constructHealthz(commit string) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		out := fmt.Sprintf(`{"commit":"%s"}`, commit)
		fmt.Fprintf(w, out)
	}
}

func findByStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status := r.URL.Query().Get("status")
	if status == "" {
		message := "Invalid status value"
		messageResponseJSON(w, http.StatusBadRequest, message)
		return
	}
	jsonResponse(w, http.StatusOK, database.FindPetByStatus(status))
}

func handleFindPet() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		petId := ps.ByName("petID")
		if petId == "findByStatus" {
			findByStatus(w, r, ps)
			return
		}
		p, err := strconv.Atoi(petId)
		if err != nil {
			message := "Invalid ID supplied"
			messageResponseJSON(w, http.StatusBadRequest, message)
			return
		}
		pet, err := database.FindPetByID(p)
		if err != nil {
			message := "Pet not found"
			messageResponseJSON(w, http.StatusNotFound, message)
			return
		}
		jsonResponse(w, http.StatusOK, pet)
	}
}

func handleDeletePet() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		petId := ps.ByName("petID")
		p, err := strconv.Atoi(petId)
		if err != nil {
			message := "Invalid ID supplied"
			messageResponseJSON(w, http.StatusBadRequest, message)
			return
		}
		err = database.DeletePet(p)
		if err != nil {
			message := "Pet not found"
			messageResponseJSON(w, http.StatusNotFound, message)
			return
		}
		messageResponseJSON(w, http.StatusOK, "Pet Deleted")
	}
}

func handleNotImplemented() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		message := "Endpoint Not Yet Implemented"
		messageResponseJSON(w, http.StatusNotImplemented, message)
		return
	}
}

func handleAddPet() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)

			message := "Invalid input. No body"
			messageResponseJSON(w, http.StatusMethodNotAllowed, message)
			return
		}
		defer r.Body.Close()

		var pet model.Pet
		err = json.Unmarshal(body, &pet)
		if err != nil {
			fmt.Println(err)
			message := "Invalid input. Cant unmarshall"
			messageResponseJSON(w, http.StatusMethodNotAllowed, message)
			return
		}

		err = database.AddPet(pet)
		if err != nil {
			fmt.Println(err)

			message := "Invalid input. Cant insert"
			messageResponseJSON(w, http.StatusMethodNotAllowed, message)
			return
		}

		jsonResponse(w, http.StatusOK, "OK")
	}
}

func messageResponseJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	msg := MessageResponse{Message: message}

	mJSON, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(mJSON))
}

func jsonResponse(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	payload, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(payload))
}
