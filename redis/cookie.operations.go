package redis

import (
	"context"
	"fmt"
	"time"
	"uuid"
)

type Session struct {
	Id   string
	Data map[string]string
}

type SessionManager struct {
	store SessionStore
}

func NewSessionManager(store SessionStore) *SessionManager {
	return &SessionManager{
		store: store,
	}
}

func (sm *SessionManager) Store() SessionStore {
	return sm.store
}

func (sm *SessionManager) Migrate(ctx context.Context, oldSession *Session) error {

	if oldSession.Id != "" {

		err := sm.store.Destroy(ctx, oldSession.Id)

		if err != nil {
			fmt.Println("Error occured while destroying the session")
		}

	}

	oldSession.Id = uuid.New().String()

	return sm.store.Save(ctx, oldSession.Id, oldSession.Data, 24*time.Hour)
}