package persistence

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/adiubaidah/rfid-syafiiyah/pkg/random"
	"github.com/adiubaidah/rfid-syafiiyah/pkg/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func clearSantriScheduleTable(t *testing.T) {
	_, err := sqlStore.db.Exec(context.Background(), "DELETE FROM santri_schedule")
	require.NoError(t, err)
}

func createRandomSantriSchedule(t *testing.T) SantriSchedule {
	startPresence := random.RandomTimeOnly()
	startTime := startPresence.Add(time.Minute * 15)
	finishTime := startTime.Add(time.Hour * 1)

	startPresencePgx := util.ConvertToPgxTime(startPresence)
	startTimePgx := util.ConvertToPgxTime(startTime)
	finishTimePgx := util.ConvertToPgxTime(finishTime)

	arg := CreateSantriScheduleParams{
		Name:          random.RandomString(10),
		Description:   pgtype.Text{Valid: false},
		StartPresence: startPresencePgx,
		StartTime:     startTimePgx,
		FinishTime:    finishTimePgx,
	}
	santriSchedule, err := testStore.CreateSantriSchedule(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, santriSchedule)

	require.Equal(t, arg.Name, santriSchedule.Name)
	require.Equal(t, arg.Description.String, santriSchedule.Description.String)
	require.Equal(t, arg.StartPresence.Microseconds, santriSchedule.StartPresence.Microseconds)
	require.Equal(t, arg.StartTime.Microseconds, santriSchedule.StartTime.Microseconds)
	require.Equal(t, arg.FinishTime.Microseconds, santriSchedule.FinishTime.Microseconds)
	return santriSchedule
}

func TestCreateSantriSchedule(t *testing.T) {
	createRandomSantriSchedule(t)
}

func TestListSantriSchedule(t *testing.T) {
	clearSantriScheduleTable(t)
	for i := 0; i < 10; i++ {
		createRandomSantriSchedule(t)
	}

	santriSchedules, err := testStore.ListSantriSchedules(context.Background(), pgtype.Time{
		Valid: false,
	})

	t.Run("list santri schedule should not error", func(t *testing.T) {
		require.NoError(t, err)
		require.Len(t, santriSchedules, 10)
	})

	t.Run("Get last santri schedule", func(t *testing.T) {
		santriSchedule, err := testStore.GetLastSantriSchedule(context.Background())
		require.NoError(t, err)
		require.NotEmpty(t, santriSchedule)
		require.Equal(t, santriSchedules[len(santriSchedules)-1].ID, santriSchedule.ID)
	})
}

func TestGetSantriSchedule(t *testing.T) {
	clearSantriScheduleTable(t)
	santriSchedule := createRandomSantriSchedule(t)
	fmt.Println(util.ConvertToTime(santriSchedule.StartTime))

	foundSantriSchedule, err := testStore.GetSantriSchedule(context.Background(), pgtype.Time{
		Microseconds: santriSchedule.StartTime.Microseconds,
		Valid:        true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, foundSantriSchedule)

	require.Equal(t, santriSchedule.ID, foundSantriSchedule.ID)
	require.Equal(t, santriSchedule.Name, foundSantriSchedule.Name)
	require.Equal(t, santriSchedule.Description.String, foundSantriSchedule.Description.String)
	require.GreaterOrEqual(t, santriSchedule.StartPresence.Microseconds, foundSantriSchedule.StartPresence.Microseconds)
	require.LessOrEqual(t, santriSchedule.StartPresence.Microseconds, foundSantriSchedule.FinishTime.Microseconds)
}

func TestUpdateSantriSchedule(t *testing.T) {
	clearSantriScheduleTable(t)
	santriSchedule := createRandomSantriSchedule(t)

	startPresence := random.RandomTimeOnly()
	startTime := startPresence.Add(time.Minute * 15)
	finishTime := startTime.Add(time.Hour * 1)
	arg := UpdateSantriScheduleParams{
		ID:            santriSchedule.ID,
		Name:          pgtype.Text{String: random.RandomString(10), Valid: true},
		Description:   pgtype.Text{Valid: false},
		StartPresence: util.ConvertToPgxTime(startPresence),
		StartTime:     util.ConvertToPgxTime(startTime),
		FinishTime:    util.ConvertToPgxTime(finishTime),
	}
	updatedSantriSchedule, err := testStore.UpdateSantriSchedule(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSantriSchedule)

	require.Equal(t, arg.ID, updatedSantriSchedule.ID)
	require.Equal(t, arg.Name.String, updatedSantriSchedule.Name)
	require.Equal(t, arg.Description.String, updatedSantriSchedule.Description.String)
	require.Equal(t, arg.StartPresence, updatedSantriSchedule.StartPresence)
	require.Equal(t, arg.StartTime, updatedSantriSchedule.StartTime)
	require.Equal(t, arg.FinishTime, updatedSantriSchedule.FinishTime)
}

func TestDeleteSantriSchedule(t *testing.T) {
	clearSantriScheduleTable(t)
	santriSchedule := createRandomSantriSchedule(t)

	deletedSantriSchedule, err := testStore.DeleteSantriSchedule(context.Background(), santriSchedule.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedSantriSchedule)

	require.Equal(t, santriSchedule.ID, deletedSantriSchedule.ID)
	require.Equal(t, santriSchedule.Name, deletedSantriSchedule.Name)
	require.Equal(t, santriSchedule.Description.String, deletedSantriSchedule.Description.String)
	require.Equal(t, santriSchedule.StartPresence, deletedSantriSchedule.StartPresence)
	require.Equal(t, santriSchedule.StartTime, deletedSantriSchedule.StartTime)
	require.Equal(t, santriSchedule.FinishTime, deletedSantriSchedule.FinishTime)
}
