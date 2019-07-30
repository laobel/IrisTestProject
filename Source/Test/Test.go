package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/couchbase/gocb.v1"
	"gopkg.in/couchbase/gocb.v1/cbft"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"
)

type Terrain struct {
	X     uint64 `json:"x"`
	Y     uint64 `json:"y"`
	Z     uint64 `json:"z"`
	Value []byte `json:"value"`
}

type Layer struct {
	Value []byte `json:"value"`
}

func saveFileToDb(filename string, bucket *gocb.Bucket) (error) {
	terrain := &Terrain{}
	layer := &Layer{}

	filenameWithSuffix := path.Base(filename)                //获取文件名带后缀
	fileSuffix := path.Ext(filenameWithSuffix)               //获取文件后缀
	filenameOnly := strings.TrimSuffix(filename, fileSuffix) //不包含文件后缀的文件路径

	filenameSlice := strings.Split(filenameOnly, "/") //分割路径
	len := len(filenameSlice)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Open", err)
		return err
	}

	fileSuffix = strings.ToUpper(fileSuffix) //字符串转大写
	//fmt.Printf("后缀: %s\n", fileSuffix)

	if fileSuffix == ".TERRAIN" {
		z := filenameSlice[len-3]
		x := filenameSlice[len-2]
		y := filenameSlice[len-1]

		terrain.Y, _ = strconv.ParseUint(y, 10, 64)
		terrain.X, _ = strconv.ParseUint(x, 10, 64)
		terrain.Z, _ = strconv.ParseUint(z, 10, 64)
		terrain.Value = b

		key := z + "-" + x + "-" + y

		_, err := bucket.Get(key, &terrain)

		if err != nil {
			_, err := bucket.Upsert(key, terrain, 0)

			if err != nil {
				time.Sleep(time.Duration(30)*time.Second) //休眠30秒
				_, err := bucket.Upsert(key, terrain, 0)
				if err != nil {
					fmt.Println("保存失败:", err)
					return err
				}
			}
		}

	} else if fileSuffix == ".JSON" {
		key := filenameSlice[len-1]

		layer.Value = b

		_, err := bucket.Get(key, &terrain)

		if err != nil {
			_, err := bucket.Upsert(key, layer, 0)

			if err != nil {
				time.Sleep(time.Duration(30)*time.Second) //休眠30秒
				_, err := bucket.Upsert(key, layer, 0)
				if err != nil {
					fmt.Println("保存失败:", err)
					return err
				}
			}
		}
	} else {
		return errors.New("文件格式不正确")
	}

	return nil
}

func saveFilesToDb(pathname string, bucket *gocb.Bucket) (error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			err := saveFilesToDb(fullDir, bucket)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			err := saveFileToDb(fullName, bucket)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return err
			}
		}
	}

	return nil
}

func simpleTextQuery(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("swanky")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query Error:", err.Error())
	}

	printResult("Simple Text Query", result)
}

func simpleTextQueryOnStoredField(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewMatchQuery("MDG").Field("destinationairport")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query on Stored Field Error:", err.Error())
	}

	printResult("Simple Text Query on Stored Field", result)
}

func printResult(label string, results gocb.SearchResults) {
	fmt.Println()
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println()
	fmt.Println(label)
	fmt.Println()

	for _, row := range results.Hits() {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}
}
func printResult1(label string, results gocb.ViewResults) {
	fmt.Println()
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println()
	fmt.Println(label)
	fmt.Println()

	var row interface{}
	for results.Next(&row) {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}
	/*for _, row := range results. {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}*/


}

func simpleTextQueryOnNonDefaultIndex(b *gocb.Bucket) {
	indexName := "travel-sample-index-hotel-description"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("swanky")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query on Non-Default Index Error:", err.Error())
	}

	printResult("Simple Text Query on Non-Default Index", result)
}
func textQueryOnStoredFieldWithFacet(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("La Rue Saint Denis!!").
		Field("reviews.content")).Limit(10).Highlight(gocb.DefaultHighlightStyle).
		AddFacet("Countries Referenced", cbft.NewTermFacet("country", 5))

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Match Query with Facet, Result by Row Error:", err.Error())
	}

	printResult("Match Query with Facet, Result by hits:", result)

	fmt.Println()
	fmt.Println("Match Query with Facet, Result by facet:")
	for _, row := range result.Facets() {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}
}

func docIdQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewDocIdQuery("hotel_26223", "hotel_28960"))

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("DocId Query Error:", err.Error())
	}

	printResult("DocId Query", result)
}

func unAnalyzedTermQuery(b *gocb.Bucket, fuzzinessLevel int) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewTermQuery("sushi").Field("reviews.content").
		Fuzziness(fuzzinessLevel)).Limit(50).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Printf("Unanalyzed Term Query with Fuzziness Level of %d Error: %s\n", fuzzinessLevel, err.Error())
	}

	printResult(fmt.Sprintf("Unanalyzed Term Query with Fuzziness Level of %d", fuzzinessLevel), result)
}

func matchPhraseQueryOnStoredField(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewMatchPhraseQuery("Eiffel Tower").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Match Phrase Query, using Analysis Error:", err.Error())
	}

	printResult("Match Phrase Query, using Analysis", result)
}

func unAnalyzedPhraseQuery(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewPhraseQuery("dorm", "rooms").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Phrase Query, without Analysis Error:", err.Error())
	}

	printResult("Phrase Query, without Analysis", result)
}

func conjunctionQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	firstQuery := cbft.NewMatchQuery("La Rue Saint Denis!!").Field("reviews.content")
	secondQuery := cbft.NewMatchQuery("boutique").Field("description")

	conjunctionQuery := cbft.NewConjunctionQuery(firstQuery, secondQuery)

	result, err := b.ExecuteSearchQuery(gocb.NewSearchQuery(indexName, conjunctionQuery).
		Limit(10).Highlight(gocb.DefaultHighlightStyle))
	if err != nil {
		fmt.Println()
		fmt.Println("Conjunction Query Error:", err.Error())
	}

	printResult("Conjunction Query", result)
}

func queryStringMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewQueryStringQuery("description: Imperial")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Query String Query Error:", err.Error())
	}

	printResult("Query String Query", result)
}

func wildCardQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewWildcardQuery("bouti*ue").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Wild Card Query Error:", err.Error())
	}

	printResult("Wild Card Query", result)
}

func numericRangeQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewNumericRangeQuery().
		Min(10100, true).Max(10200, true).Field("id")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Wild Card Query Error:", err.Error())
	}

	printResult("Wild Card Query", result)
}

func regexpQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewRegexpQuery("[a-z]").
		Field("description")).Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Regexp Query Error:", err.Error())
	}

	printResult("Regexp Query", result)
}


func geoQueryTest(b *gocb.Bucket) {
	indexName := "mygeoindex"
	query := gocb.NewSearchQuery(indexName, cbft.NewGeoDistanceQuery(53.482358,-2.235143,"100mi")).Sort(cbft.NewSearchSortGeoDistance("name",53.482358,-2.235143))

	//query := gocb.NewSearchQuery(indexName, cbft.NewSearchSortGeoDistance("name",53.482358,-2.235143))
	//query := gocb.NewSearchQuery(indexName, cbft.NewGeoBoundingBoxQuery(53.482358,-2.235143,40.991862,28.955043)).Sort("name")

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Geo Query Error:", err.Error())
	}

	printResult("Geo Query", result)
}


func spatialQueryTest(b *gocb.Bucket){
	query := gocb.NewSpatialQuery("city","cityView")

	bounds:=[]float64{102,24,103,25}
	query.Bbox(bounds)

	result, err := b.ExecuteSpatialQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Geo Query Error:", err.Error())
	}

	//printResult1("Geo Query", string(result.NextBytes()))
	printResult1("Geo Query", result)
}

func main() {
	cluster, _ := gocb.Connect("couchbase://192.168.1.10")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "123369",
	})
	bucket, _ := cluster.OpenBucket("TestBucket", "")

	/*simpleTextQuery(bucket)
	simpleTextQueryOnStoredField(bucket)
	simpleTextQueryOnNonDefaultIndex(bucket)
	textQueryOnStoredFieldWithFacet(bucket)
	docIdQueryMethod(bucket)
	unAnalyzedTermQuery(bucket, 0)
	unAnalyzedTermQuery(bucket, 2)
	matchPhraseQueryOnStoredField(bucket)
	unAnalyzedPhraseQuery(bucket)
	conjunctionQueryMethod(bucket)
	queryStringMethod(bucket)
	wildCardQueryMethod(bucket)
	numericRangeQueryMethod(bucket)
	regexpQueryMethod(bucket)
*/

	//geoQueryTest(bucket)
	bucket, _ = cluster.OpenBucket("yunnan-city", "")
	spatialQueryTest(bucket)
	//bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	//saveFilesToDb("D:/apache-tomcat-9.0.14/webapps/tiles",bucket)

	//saveFileToDb("D:/apache-tomcat-9.0.14/webapps/result/0/1/0.terrain",bucket)
	//saveFileToDb("D:/apache-tomcat-9.0.14/webapps/result/layer.json", bucket)

	/*var terrain Terrain
	_,err:=bucket.Get("010", &terrain)
	fmt.Printf("地形: %v\n", err)

	var layer Layer
	bucket.Get("layer", &layer)
	fmt.Printf("Layer: %v\n", string(layer.Value))*/

	/*bucket.Upsert("u:kingarthur",
		User{
			Id:        "kingarthur",
			Email:     "kingarthur@couchbase.com",
			Interests: []string{"Holy Grail", "African Swallows"},
		}, 0)

	// Get the value back
	var inUser User
	bucket.Get("u:kingarthur", &inUser)
	fmt.Printf("User: %v\n", inUser)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM bucketname WHERE $1 IN interests")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"African Swallows"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}*/
}
