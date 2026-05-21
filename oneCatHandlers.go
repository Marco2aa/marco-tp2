package main

import (
	"net/http"
)

func getCat(req *http.Request) (int, any) {
	catID := req.PathValue("catId")

	cat, found := catsDatabase[catID]
	if !found {
		Logger.Infof("Cat '%s' not found", catID)
		return http.StatusNotFound, "Cat not found"
	}

	if cat.ID == "" {
		cat.ID = catID
	}

	Logger.Infof("Cat '%s' found", catID)
	return http.StatusOK, cat
}

func deleteCat(req *http.Request) (int, any) {
	catID := req.PathValue("catId")

	if _, found := catsDatabase[catID]; !found {
		Logger.Infof("Cat '%s' not found", catID)
		return http.StatusNotFound, "Cat not found"
	}

	delete(catsDatabase, catID)

	Logger.Infof("Cat '%s' deleted", catID)
	return http.StatusNoContent, nil
}