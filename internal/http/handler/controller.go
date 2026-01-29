package handler

import (
	"encoding/json"
	"expensemanagement/internal/domain/request"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"expensemanagement/internal/domain/resources/resourceserrormessagesrepository"
	"expensemanagement/internal/services"
	"expensemanagement/internal/validation"
	"net/http"
)

type ExpensesServices struct {
	Services services.Expenses
}

func (s *ExpensesServices) RegisterExpenses(w http.ResponseWriter, r *http.Request) {
	var expenses request.Expenses

	if err := json.NewDecoder(r.Body).Decode(&expenses); err != nil {
		http.Error(w, resourceserrormessages.CouldNotDecodeJson, http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	validationErrors := validation.ListErrorMessages(expenses)

	if len(validationErrors) > 0 {
		msg, _ := json.MarshalIndent(validationErrors, "", "  ")
		http.Error(w, string(msg), http.StatusBadRequest)
		return
	}

	if err := s.Services.Create(expenses); err != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotInsertIntoDatabase, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
