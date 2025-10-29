package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/maaviah17/students-api/internal/types"
	"github.com/maaviah17/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		slog.Info("creating a student")
		var newStudent types.Student

		//whatever data that is coming in will be json decoded and stored in the struct.
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty bodyy aaayi hai :(( ")))
			return 
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		//request validation
		if err := validator.New().Struct(newStudent); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return 
		}

		response.WriteJson(w, http.StatusCreated, map[string]string {"success":"OK"})
	} 
}