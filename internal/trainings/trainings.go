package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// Разделяем строку datastring на слайс строк
	parts := strings.Split(datastring, ",")
	
	// Проверяем, что получилось ровно 3 элемента
	if len(parts) != 3 {
		return errors.New("invalid string format: 3 fields are expected (steps, type, duration)")
	}

	// Преобразуем первый элемент слайса (количество шагов) в тип int
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("failed to convert the number of steps to a number: " + stepsStr)
	}
	if steps <= 0 {
		return errors.New("the number of steps must be greater than zero")
	}
	t.Steps = steps

	// Сохраняем значение типа тренировки
	t.TrainingType = parts[1]

	// Преобразуем третий элемент слайса в time.Duration
	durationStr := parts[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("couldn't convert duration to format time.Duration: " + durationStr)
	}
	if duration <= 0 {
		return errors.New("the duration of the workout must be greater than zero")
	}	
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Проверяем, что количество шагов не отрицательное
	if t.Steps < 0 {
		return "", errors.New("the number of steps cannot be negative")
	}
	if t.Personal.Weight <= 0 {
		return "", errors.New("the weight must be a positive number")
	}
	if t.Personal.Height <= 0 {
		return "", errors.New("the height must be a positive number")
	}
	if t.Duration <= 0 {
		return "", errors.New("the duration of the workout should be positive")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	
	// Вычисляем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	
	// Рассчитываем калории в зависимости от типа тренировки
	var calories float64
	var err error

	switch t.TrainingType {
	case "бег", "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "ходьба", "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("unknown type of training")
	}
	
	// Формируем итоговую строку с результатами
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
