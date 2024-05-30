package player

import (
	"fmt"
	"github/cesardev31/go-clean/internal/domain"
	"log"
	"time"
)

func (s Service) Create(player domain.Player) (id interface{}, err error) {

	player.CreationTime = time.Now().UTC()

	insertedID, err := s.Repo.Insert(player)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error Creating player: %w", err)
	}

	return insertedID, nil
}
