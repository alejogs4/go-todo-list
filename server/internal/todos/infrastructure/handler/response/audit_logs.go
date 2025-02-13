package response

import (
	"time"

	"go-todo-list.com/m/internal/todos/infrastructure/persistance"
)

type TodoAuditLogResponse struct {
	ID          int    `json:"id"`
	TodoID      string `json:"todo_id"`
	Action      string `json:"action"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}

func FromAuditsToResponse(audits []persistance.TodoAudit) []TodoAuditLogResponse {
	response := make([]TodoAuditLogResponse, 0, len(audits))

	for _, audit := range audits {
		response = append(response, FromAuditToResponse(audit))
	}

	return response
}

func FromAuditToResponse(auditLog persistance.TodoAudit) TodoAuditLogResponse {
	return TodoAuditLogResponse{
		ID:          auditLog.ID,
		TodoID:      auditLog.TodoID,
		Action:      auditLog.Action,
		CreatedAt:   auditLog.CreatedAt.Format(time.DateTime),
		Description: auditLog.Description,
	}
}
