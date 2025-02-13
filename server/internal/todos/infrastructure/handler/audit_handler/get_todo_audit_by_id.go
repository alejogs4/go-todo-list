package audit_handler

import (
	"errors"
	"net/http"
	"strconv"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/response"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func HandleGetTodoAuditByID(auditRepository persistance.AuditRepository) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := api.GetPathParam(request, "id")
		if err != nil {
			api.InvalidRequest(writer, "Missing audit id")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			api.InvalidRequest(writer, "Invalid audit id it should be a number")
			return
		}

		auditLog, err := auditRepository.GetByID(request.Context(), idInt)
		if err != nil {
			if errors.Is(err, persistance.ErrNotFoundAudit) {
				api.NotFound(writer, "Audit log not found")
				return
			}

			api.InternalServerError(writer, err.Error())
			return
		}

		api.Success(writer, response.FromAuditToResponse(auditLog))
	}
}
