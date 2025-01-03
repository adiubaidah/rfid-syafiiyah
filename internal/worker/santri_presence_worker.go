package worker

import (
	"context"
	"time"

	"github.com/adiubaidah/rfid-syafiiyah/internal/constant/model"
	db "github.com/adiubaidah/rfid-syafiiyah/internal/storage/persistence"
	"github.com/adiubaidah/rfid-syafiiyah/internal/usecase"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type SantriPresenceWorker interface {
	CreateAlphaForMissingPresence(scheduleId int32)
}

type santriPresenceWorker struct {
	logger  *logrus.Logger
	usecase usecase.SantriPresenceUseCase
}

func NewSantriPresenceWorker(logger *logrus.Logger, usecase usecase.SantriPresenceUseCase) SantriPresenceWorker {
	return &santriPresenceWorker{
		logger:  logger,
		usecase: usecase,
	}
}

func (s *santriPresenceWorker) CreateAlphaForMissingPresence(scheduleId int32) {
	missing, err := s.usecase.ListMissingSantriPresences(context.Background(), &model.ListMissingSantriPresenceRequest{
		ScheduleID: scheduleId,
		Time:       time.Now(),
	})

	s.logger.Println("missing :", missing)

	if err != nil {
		return
	}

	var bulkPresenceArgs []db.CreateSantriPresencesParams

	for _, m := range *missing {
		bulkPresenceArgs = append(bulkPresenceArgs, db.CreateSantriPresencesParams{
			ScheduleID:         scheduleId,
			SantriID:           m.Id,
			Type:               db.PresenceTypeAlpha,
			Notes:              pgtype.Text{String: "Auto created by system", Valid: true},
			CreatedBy:          db.PresenceCreatedByTypeSystem,
			SantriPermissionID: pgtype.Int4{Int32: 0, Valid: false},
			CreatedAt:          pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
	}

	_, err = s.usecase.BulkCreateSantriPresence(context.Background(), bulkPresenceArgs)
	if err != nil {
		s.logger.Error(err)
		return
	}

}
