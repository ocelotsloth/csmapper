package csmapper

import (
	"log"
	"strconv"
	"time"

	"github.com/ocelotsloth/goqrz"
	geojson "github.com/paulmach/go.geojson"
)

// Club is an ARRL club, provided via CSV file.
type Club struct {
	Name      string `csv:"Club"`   // Name of the ARRL club
	Callsign  string `csv:"Call"`   // Callsign
	County    string `csv:"County"` // City/County
	Latitude  string `csv:"Lat"`    // GPS Latitude
	Longitude string `csv:"Log"`    // GPS Longitude
}

const (
	GOQRZ_USER_AGENT = "gh/ocelotsloth/csmapper" // User agent to report to QRZ.com
	GOQRZ_TPS        = 3                         // Max TPS to use with QRZ.com (be nice to their endpoint)
)

func GenerateGeoJSON(user string, pass string, clubs []Club) *geojson.FeatureCollection {
	log.Println("getting session key")
	sessionKey, err := goqrz.GetSessionKey(user, pass, GOQRZ_USER_AGENT)
	if err != nil {
		return nil
	}
	log.Println("got session key")
	log.Println("getting session")
	qrzSession, err := goqrz.GetSession(sessionKey, GOQRZ_USER_AGENT)
	if err != nil {
		return nil
	}
	log.Println("got session")

	features := geojson.NewFeatureCollection()

	log.Println("processing clubs")
	for _, club := range clubs {
		log.Println("start new club, calling qrz")
		callsign, err := goqrz.GetCallsign(qrzSession.Key, club.Callsign, GOQRZ_USER_AGENT)
		if err != nil {
			return nil
		}
		log.Println("got callsign from qrz")
		//spew.Dump(callsign)
		// sleep for inverse TPS
		log.Println("starting tps sleep")
		time.Sleep((1 / GOQRZ_TPS) * time.Second)
		log.Println("finished tps sleep")

		log.Println("generating lat/long floats")
		lat, err := strconv.ParseFloat(callsign.Lat, 64)
		if err != nil {
			return nil
		}
		lon, err := strconv.ParseFloat(callsign.Lon, 64)
		if err != nil {
			return nil
		}

		log.Println("generating point feature")
		clubPoint := geojson.NewPointFeature([]float64{lon, lat})
		clubPoint.SetProperty("title", club.Name)
		clubPoint.SetProperty("description", club.County)
		log.Println("adding point feature")
		features.AddFeature(clubPoint)
	}
	log.Println("returning feature set")
	return features
}
