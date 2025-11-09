package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверяем корректность входных параметров
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be negative")
	}
	if weight <= 0 {
		return 0, errors.New("the weight must be a positive number")
	}
	if height <= 0 {
		return 0, errors.New("the height must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("the duration of the run should be positive")
	}
	
	// Рассчитываем среднюю скорость с помощью существующей функции
	meanSpeed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитываем количество потраченных калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// Применяем корректирующий коэффициент для ходьбы
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//Проверяем корректность входных данных
	if steps <= 0 {
		return 0, errors.New("the number of steps cannot be negative")
	}
	if weight <= 0 {
		return 0, errors.New("the weight must be a positive number")
	}
	if height <= 0 {
		return 0, errors.New("the height must be a positive number")
	}
	if duration <= 0 {
		return 0, errors.New("the duration of the run should be positive")
	}

	//Рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)
	
	//Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()
	
	//Рассчитываем количество потраченных калорий
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	//Проверяем отрицательные шаги
	if steps < 0 {
		return 0
	}
		
	//Проверяем отрицательную продолжительность
	if duration <= 0 {
		return 0
	}

	//Вычисляем дистанцию
	distance := Distance(steps, height)

	//Переводим продолжительность в часы
	durationHours := duration.Hours()

	//Вычисляем среднюю скорость
	meanSpeed := distance / durationHours

	return meanSpeed	
}

func Distance(steps int, height float64) float64 {
	//Рассчитываем длину шага
	stepLength := height * stepLengthCoefficient
	
	//Умножаем пройденное количество шагов на длину шага
	totalDistanceMeters := float64(steps) * stepLength
	
	//Разделяем полученное значение на число метров в километре
	distanceKilometers := totalDistanceMeters / mInKm
	
	return distanceKilometers
}
