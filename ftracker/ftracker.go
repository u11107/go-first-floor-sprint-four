package ftracker

import (
	"fmt"
	"math"
)

const (
	lenStep   = 0.65
	mInKm     = 1000
	minInH    = 60
	kmhInMsec = 0.278
	cmInM     = 100
)

func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch {
	case trainingType == "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч."+
			"\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories)

	case trainingType == "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч."+
			"\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories)

	case trainingType == "Плавание":
		distance := distance(action)
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч."+
			"\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

const (
	runningCaloriesMeanSpeedMultiplier = 18
	runningCaloriesMeanSpeedShift      = 1.79
)

func RunningSpentCalories(action int, weight, duration float64) float64 {

	return (runningCaloriesMeanSpeedMultiplier * meanSpeed(action,
		duration) * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH
}

const (
	walkingCaloriesWeightMultiplier = 0.035
	walkingSpeedHeightMultiplier    = 0.029
)

func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	// ваш код здесь
	speedByMInSec := meanSpeed(action, duration) * kmhInMsec
	heightInM := height / cmInM
	return (walkingCaloriesWeightMultiplier*weight + (math.Pow(speedByMInSec,
		2)/heightInM)*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

const (
	swimmingCaloriesMeanSpeedShift   = 1.1
	swimmingCaloriesWeightMultiplier = 2
)

func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {

	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
