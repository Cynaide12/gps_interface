package lib_geofence

import (
	"math"
)

// IsInsideGeofence проверяет, находится ли точка внутри геозоны
// Возвращает true если расстояние до центра меньше или равно радиусу
func IsInsideGeofence(lat, lon, centerLat, centerLon, radius float64) bool {
	distance := Haversine(lat, lon, centerLat, centerLon)
	return distance <= radius
}

// Haversine рассчитывает расстояние между точками в метрах
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Радиус Земли в метрах
	
	// Переводим градусы в радианы
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	// Формула гаверсинусов
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}