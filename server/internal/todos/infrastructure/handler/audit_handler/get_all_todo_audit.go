package audit_handler

import (
	"errors"
	"net/http"

	"go-todo-list.com/m/internal/shared/api"
	"go-todo-list.com/m/internal/todos/infrastructure/handler/response"
	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

func HandleGetAllTodoAudit(auditRepository persistance.AuditRepository) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		auditsLogs, err := auditRepository.GetAll(request.Context())
		if err != nil {
			if errors.Is(err, persistance.ErrNotFoundAudit) {
				api.Success(writer, []string{})
				return
			}

			api.InternalServerError(writer, err.Error())
			return
		}

		api.Success(writer, response.FromAuditsToResponse(auditsLogs))
	}
}
