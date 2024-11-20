package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (store *SQLStore) CreateArduinoWithModes(ctx context.Context, arduinoName string, modeParams []CreateArduinoModesParams) (Arduino, error) {
	var createdArduino Arduino

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		arduino, err := q.CreateArduino(ctx, arduinoName)
		if err != nil {
			return err
		}
		createdArduino = arduino

		for i := range modeParams {
			modeParams[i].ArduinoID = arduino.ID
		}
		_, err = q.CreateArduinoModes(ctx, modeParams)
		return err

	})
	return createdArduino, err
}

func (store *SQLStore) UpdateArduinoWithModes(ctx context.Context, arduinoID int32,
	name string,
	modeParams []CreateArduinoModesParams) (Arduino, error) {
	var updatedArduino Arduino

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		arduino, err := q.UpdateArduino(ctx, UpdateArduinoParams{
			ID:   arduinoID,
			Name: pgtype.Text{String: name, Valid: name != ""},
		})
		if err != nil {
			return err
		}

		updatedArduino = arduino

		err = q.DeleteArduinoModeByArduinoId(ctx, arduinoID)
		if err != nil {
			return err
		}

		for i := range modeParams {
			modeParams[i].ArduinoID = arduinoID
		}

		_, err = q.CreateArduinoModes(ctx, modeParams)
		return err

	})
	return updatedArduino, err
}

func (store *SQLStore) CreateHolidayWithDates(ctx context.Context, arg CreateHolidayParams, argsCreateDates []CreateHolidayDatesParams) (Holiday, error) {
	var createdHoliday Holiday

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		holiday, err := q.CreateHoliday(ctx, arg)
		if err != nil {
			return err
		}
		createdHoliday = holiday
		var args []CreateHolidayDatesParams
		for _, arg := range argsCreateDates {
			arg.HolidayID = holiday.ID
			args = append(args, arg)

		}
		_, err = q.CreateHolidayDates(ctx, args)

		return err
	})
	return createdHoliday, err
}
func (store *SQLStore) UpdateHolidayWithDates(ctx context.Context, holidayId int32, arg UpdateHolidayParams, argsCreateDates []CreateHolidayDatesParams) (Holiday, error) {
	var updateHoliday Holiday

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		holiday, err := q.UpdateHoliday(ctx, arg)
		if err != nil {
			return err
		}
		updateHoliday = holiday

		err = q.DeleteHolidayDateByHolidayId(ctx, holidayId)
		if err != nil {
			return err
		}

		var args []CreateHolidayDatesParams
		for _, arg := range argsCreateDates {
			arg.HolidayID = holiday.ID
			args = append(args, arg)

		}
		_, err = q.CreateHolidayDates(ctx, args)

		return err
	})
	return updateHoliday, err
}