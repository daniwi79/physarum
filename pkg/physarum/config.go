package physarum

import (
	"math/rand"
)

type Config struct {
	PopulationPercentage float64
	SensorAngle          float64
	SensorDistance       float64
	RotationAngle        float64
	StepDistance         float64
	DepositionAmount     float64
	DecayFactor          float64
}

func RandomConfig() *Config {
	sensorAngle := rand.Float64() * Radians(120)
	sensorDistance := rand.Float64() * 64
	rotationAngle := rand.Float64() * Radians(120)
	stepDistance := 1 + rand.NormFloat64()*0.25
	decayFactor := 0.1 + rand.NormFloat64()*0.01
	return &Config{
		PopulationPercentage: 0.5,
		SensorAngle:          sensorAngle,
		SensorDistance:       sensorDistance,
		RotationAngle:        rotationAngle,
		StepDistance:         stepDistance,
		DepositionAmount:     5,
		DecayFactor:          decayFactor,
	}
}

func RandomConfigs(n int) []*Config {
	configs := make([]*Config, n)
	for i := range configs {
		configs[i] = RandomConfig()
	}
	return configs
}