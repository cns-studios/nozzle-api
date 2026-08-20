package github_bot

import (
	"fmt"
	"net/http"
	"nozzle-api/utils"
	"os"

	"github.com/gorilla/mux"
)

func GetAllrepos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]

	requestURL := fmt.Sprintf("https://api.github.com/user/%s/repos", userid)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	response := utils.ReadJSONResponse(res)

	var repoNames []string

	data, ok := response["data"].([]any)
	if !ok {
		return
	}

	username := getUsername(userid)

	for _, item := range data {
		if repoMap, ok := item.(map[string]any); ok {
			if name, ok := repoMap["name"].(string); ok {
				repoNames = append(repoNames, username+"/"+name)
			}
		}
	}
	utils.WriteJSON(w, http.StatusOK, repoNames)
}

func getUsername(id string) string {
	requestURL := fmt.Sprintf("https://api.github.com/user/%s", id)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	response := utils.ReadJSONResponse(res)

	return response["login"].(string)
}
