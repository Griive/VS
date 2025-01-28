package hendlers

import "context"

type storage interface {
	GetUserById(id string)
}

func (uc *Service) GetUserById(ctx context.Context, id string) {
	return s.storage
}
