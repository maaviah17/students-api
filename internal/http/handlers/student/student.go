package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/maaviah17/students-api/internal/types"
	"github.com/maaviah17/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		var newStudent types.Student

		//whatever data that is coming in will be json decoded and stored in the struct.
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return 
		}
		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string {"success":"OK"})
	} 
}