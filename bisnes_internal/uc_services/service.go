package ucservices

import "context"

type Service interface {
	GetUserById(ctx context.Context, id string)
	GetAllTask(ctx context.Context, priority int)
}
