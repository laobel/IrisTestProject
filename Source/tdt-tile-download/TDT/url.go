package TDT

import "strings"

func Url(mapProjection, layer,serviceMode string) string {
	mapProjection = strings.ToLower(mapProjection)
	layer = strings.ToLower(layer)
	var projectionFlag string

	switch mapProjection {
	case "geographic":
		projectionFlag = "c"
	case "webmercator":
		projectionFlag = "w"
	}

	var url string
	switch serviceMode {
	case "wmts":
		url="http://%s.tianditu.gov.cn/"+layer+"_"+projectionFlag+"/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER="+layer+"&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILECOL=%d&TILEROW=%d&TILEMATRIX=%d&tk=%s"
	case "tms":
		url="http://%s.tianditu.gov.cn/DataServer?T="+layer+"_"+projectionFlag+"&x=%d&y=%d&l=%d&tk=%s"
	}

	return url
}
