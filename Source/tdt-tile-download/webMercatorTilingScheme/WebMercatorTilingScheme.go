package webMercatorTilingScheme

import (
	"IrisTestProject/Source/tdt-tile-download/ellipsoid"
	"IrisTestProject/Source/tdt-tile-download/webMercatorProjection"
)

type WebMercatorTilingScheme struct {
	ellipsoid ellipsoid.Ellipsoid
	projection webMercatorProjection.WebMercatorProjection
	numberOfLevelZeroTilesX  int
	numberOfLevelZeroTilesY  int
}

func NewWebMercatorTilingScheme(ellipsoid ellipsoid.Ellipsoid,projection webMercatorProjection.WebMercatorProjection,numberOfLevelZeroTilesX,numberOfLevelZeroTilesY int) *WebMercatorTilingScheme{
	return &WebMercatorTilingScheme{
		ellipsoid: ellipsoid,
		projection: projection,
		numberOfLevelZeroTilesX:numberOfLevelZeroTilesX,
		numberOfLevelZeroTilesY:numberOfLevelZeroTilesY,
	}
}

func (wts *WebMercatorTilingScheme) rectangleNortheastInMeters() (float64, float64) {
	return wts.ellipsoid.SemimajorAxisTimesPi(), wts.ellipsoid.SemimajorAxisTimesPi()
}

func (wts *WebMercatorTilingScheme) rectangleSouthwestInMeters() (float64, float64) {
	return -wts.ellipsoid.SemimajorAxisTimesPi(), -wts.ellipsoid.SemimajorAxisTimesPi()
}

func (wts *WebMercatorTilingScheme) NumberOfXTilesAtLevel(level int) float64{
	return float64(uint(wts.numberOfLevelZeroTilesX) << uint(level))
}

func (wts *WebMercatorTilingScheme) NumberOfYTilesAtLevel(level int) float64{
	return float64(uint(wts.numberOfLevelZeroTilesY) << uint(level))
}

func (wts *WebMercatorTilingScheme) LonLatToTileXY(lon,lat float64,zoom int)(int,int){
	var xTiles = wts.NumberOfXTilesAtLevel(zoom)
	var yTiles = wts.NumberOfYTilesAtLevel(zoom)

	rectangleNortheastInMetersX,rectangleNortheastInMetersY:=wts.rectangleNortheastInMeters()
	rectangleSouthwestInMetersX,rectangleSouthwestInMetersY:=wts.rectangleSouthwestInMeters()

	var overallWidth = rectangleNortheastInMetersX - rectangleSouthwestInMetersX
	var xTileWidth = overallWidth / xTiles
	var overallHeight = rectangleNortheastInMetersY - rectangleSouthwestInMetersY
	var yTileHeight = overallHeight / yTiles

	webMercatorPositionX, webMercatorPositionY:= wts.projection.Project(lon,lat)
	var distanceFromWest = webMercatorPositionX - rectangleSouthwestInMetersX
	var distanceFromNorth = rectangleNortheastInMetersY - webMercatorPositionY

	var xTileCoordinate = int(distanceFromWest / xTileWidth) | 0.0
	if xTileCoordinate >= int(xTiles) {
		xTileCoordinate = int(xTiles) - 1
	}
	var yTileCoordinate = int(distanceFromNorth / yTileHeight) | 0
	if yTileCoordinate >= int(yTiles) {
		yTileCoordinate = int(yTiles) - 1
	}

	return xTileCoordinate,yTileCoordinate
}
