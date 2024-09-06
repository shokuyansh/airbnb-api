package utils

import (
	"errors"
	"math/rand"
	"time"
)

// Create a new random generator with a unique seed
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Mock function: In reality, you'd use an external API or web scraping
func FetchAirbnbData(roomId string) (float64, float64, float64, float64, error) {
   
    if roomId == "" {
        return 0, 0, 0, 0, errors.New("invalid room ID")
    }

   
    occupancy := mockOccupancyPercentage()
    avgRate := mockAverageNightRate()
    highRate := mockHighestNightRate()
    lowRate := mockLowestNightRate()

    return occupancy, avgRate, highRate, lowRate, nil
}

func mockOccupancyPercentage() float64 {
   
    return 60 + seededRand.Float64()*(90-60)
}

func mockAverageNightRate() float64 {
    return 100 + seededRand.Float64()*(500-100)
}

func mockHighestNightRate() float64 {
    return 400 + seededRand.Float64()*(800-400)
}

func mockLowestNightRate() float64 {
    return 80 + seededRand.Float64()*(200-80)
}
