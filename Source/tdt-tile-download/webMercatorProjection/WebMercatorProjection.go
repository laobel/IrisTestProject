package webMercatorProjection

import (
	"IrisTestProject/Source/tdt-tile-download/ellipsoid"
	"IrisTestProject/Source/tdt-tile-download/util"
	"math"
)

var maximumLatitude=math.Pi/2.0 - (2.0 * math.Atan(math.Exp(-math.Pi)))

type WebMercatorProjection struct {
	ellipsoid ellipsoid.Ellipsoid
}

func NewWWebMercatorProjection(ellipsoid ellipsoid.Ellipsoid) *WebMercatorProjection{
	return &WebMercatorProjection{
		ellipsoid: ellipsoid,
	}
}

//lon,lat都是经纬度
func (wmp *WebMercatorProjection) Project(lon,lat float64)(float64,float642 float64){
	longitude:=util.Radians(lon)
	latitude:=util.Radians(lat)

	var semimajorAxis = wmp.ellipsoid.MaximumRadius()

	var x = longitude * semimajorAxis
	var y = wmp.GeodeticLatitudeToMercatorAngle(latitude) * semimajorAxis

	return x,y
}

//latitude是弧度
func(wmp *WebMercatorProjection) GeodeticLatitudeToMercatorAngle(latitude float64)float64{
	if latitude > maximumLatitude {
		latitude = maximumLatitude
	} else if latitude < -maximumLatitude {
		latitude = -maximumLatitude
	}
	var sinLatitude = math.Sin(latitude)

	return 0.5 * math.Log((1.0 + sinLatitude) / (1.0 - sinLatitude))
}