package tui

import (
	"context"
	"strings"
)

func (s StateTUI) viewUsersTable(ctx context.Context) string {
	users, err := s.UserUseCase.GetUsers(ctx)
	if err != nil {
		return err.Error()
	}

	var b strings.Builder
	for _, user := range users {
		b.WriteString(user.ID)
		b.WriteString(" | ")
		b.WriteString(user.Email)
		b.WriteString("\n")
	}
	return b.String()
}
