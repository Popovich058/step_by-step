package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds DaySteps) Weight() float64 {
	return ds.Personal.Weight
}

func (ds DaySteps) Height() float64 {
	return ds.Personal.Height
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем строку на слайс строк
	parts := strings.Split(datastring, ",")
	
	// Проверяем, что получилось ровно 2 элемента
	if len(parts) != 2 {
		return errors.New("invalid string format: 2 fields are expected (steps, duration)")
	}

	// Преобразуем первый элемент слайса в тип int
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("failed to convert the number of steps to a number: " + stepsStr)
	}
	if steps <= 0 {
		return errors.New("the number of steps cannot be negative")
	}
	ds.Steps = steps

	// Преобразуем второй элемент слайса в time.Duration
	durationStr := parts[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("couldn't convert duration to format time.Duration: " + durationStr)
	}
	if duration <= 0 {
		return errors.New("the duration must be a positive value.")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Проверяем корректность входных данных
	if ds.Steps < 0 {
		return "", errors.New("the number of steps cannot be negative")
	}
	if ds.Personal.Weight <= 0 {
		return "", errors.New("the weight must be a positive number.")
	}
	if ds.Personal.Height <= 0 {
		return "", errors.New("the weight must be a positive number.")
	}
	if ds.Duration <= 0 {
		return "", errors.New("the duration must be a positive value.")
	}

	// Вычисляем дистанцию 
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	// Рассчитываем сожжённые калории 
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight(),
		ds.Height(),	
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	// Формируем итоговую строку
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
