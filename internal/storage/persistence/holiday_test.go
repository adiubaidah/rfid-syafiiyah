package persistence

import (
	"context"
	"testing"

	"github.com/adiubaidah/rfid-syafiiyah/pkg/random"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func clearHoliday(t *testing.T) {
	_, err := sqlStore.db.Exec(context.Background(), `DELETE FROM "employee_occupation"`)
	require.NoError(t, err)
}

func createRandomHoliday(t *testing.T) Holiday {
	arg := CreateHolidayParams{
		Name:        random.RandomString(8),
		Color:       pgtype.Text{String: random.RandomString(7), Valid: true},
		Description: pgtype.Text{String: random.RandomString(50), Valid: true},
	}

	holiday, err := testStore.CreateHoliday(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, holiday)

	require.Equal(t, arg.Name, holiday.Name)
	require.Equal(t, arg.Description.String, holiday.Description.String)
	return holiday
}

func TestCreateHoliday(t *testing.T) {
	clearHoliday(t)
	createRandomHoliday(t)
}

func TestUpdateHoliday(t *testing.T) {
	clearHoliday(t)
	holiday := createRandomHoliday(t)

	arg := UpdateHolidayParams{
		ID:          holiday.ID,
		Name:        random.RandomString(8),
		Description: pgtype.Text{String: random.RandomString(50), Valid: true},
		Color:       pgtype.Text{String: random.RandomString(7), Valid: true},
	}

	updatedHoliday, err := testStore.UpdateHoliday(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, holiday)

	require.Equal(t, arg.Name, updatedHoliday.Name)
	require.Equal(t, arg.Description.String, updatedHoliday.Description.String)
	require.Equal(t, arg.Color.String, updatedHoliday.Color.String)
}

func TestDeleteHoliday(t *testing.T) {
	clearHoliday(t)
	holiday := createRandomHoliday(t)
	deletedHoliday, err := testStore.DeleteHoliday(context.Background(), holiday.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deletedHoliday)
}
