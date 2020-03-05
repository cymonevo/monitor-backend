package handler

import "github.com/cymonevo/monitor-backend/internal/router"

type BaseHandler interface {
	Register() router.Router
}
