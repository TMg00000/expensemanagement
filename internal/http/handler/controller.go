package handler

import (
	"encoding/json"
	"expensemanagement/internal/domain/request"
	"expensemanagement/internal/domain/resources/resourceserrormessages"
	"expensemanagement/internal/domain/resources/resourceserrormessagesrepository"
	"expensemanagement/internal/services"
	"expensemanagement/internal/validation"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpensesServices struct {
	Services services.Expenses
}

func (s *ExpensesServices) RegisterExpenses(w http.ResponseWriter, r *http.Request) {
	var insertExpenses request.Expenses

	if err := json.NewDecoder(r.Body).Decode(&insertExpenses); err != nil {
		http.Error(w, resourceserrormessages.CouldNotDecodeJson, http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	validationErrors := validation.ListErrorMessages(insertExpenses)

	if len(validationErrors) > 0 {
		msg, _ := json.MarshalIndent(validationErrors, "", "  ")
		http.Error(w, string(msg), http.StatusBadRequest)
		return
	}

	if err := s.Services.Create(insertExpenses); err != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotInsertIntoDatabase, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertExpenses)
}

func (s *ExpensesServices) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	AllExpenses, err := s.Services.GetAll()
	if err != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotFindAnyExpense, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AllExpenses)
}

func (s *ExpensesServices) UpdateExpenses(w http.ResponseWriter, r *http.Request) {
	var updateExpenses request.Expenses

	if err := json.NewDecoder(r.Body).Decode(&updateExpenses); err != nil {
		http.Error(w, resourceserrormessages.CouldNotDecodeJson, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, errId := primitive.ObjectIDFromHex(idStr)
	if errId != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotFindId, http.StatusBadRequest)
		return
	}
	updateExpenses.Id = id

	validationErrors := validation.ListErrorMessages(updateExpenses)

	if len(validationErrors) > 0 {
		msg, _ := json.MarshalIndent(validationErrors, "", "  ")
		http.Error(w, string(msg), http.StatusBadRequest)
		return
	}

	if err := s.Services.UpdateById(updateExpenses); err != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotInsertIntoDatabase, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateExpenses)
}

func (s *ExpensesServices) DeleteExpensesById(w http.ResponseWriter, r *http.Request) {
	var deleteById request.Expenses

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, errId := primitive.ObjectIDFromHex(idStr)
	if errId != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotFindId, http.StatusBadRequest)
		return
	}
	deleteById.Id = id

	if err := s.Services.DeleteById(deleteById); err != nil {
		http.Error(w, resourceserrormessagesrepository.CouldNotDeleteExpense, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *ExpensesServices) DeleteAllExpenses(w http.ResponseWriter, r *http.Request) {
	var deleteAll request.Expenses

	if err := s.Services.DeleteAll(deleteAll); err != nil {
		http.Error(w, resourceserrormessagesrepository.ThereIsNoExpenseToDelete, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
