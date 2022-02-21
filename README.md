# callsign-mapper

Utility to map Amateur Radio callsigns.

Takes a CSV file of clubs, queries QRZ for updated coordinates, and produces a GeoJSON output for mapping.

## Usage

```
NAME:
   csmapper geojson - Generate a geojson file from callsign list

USAGE:
   csmapper geojson [command options] [arguments...]

OPTIONS:
   --user value, -u value   QRZ.com Username
   --pass value, -p value   QRZ.com Password
   --clubs value, -c value  Clubs CSV File
```

## Example

### Clubs CSV

```csv
Club,Call,County,Lat,Log
Albemarle ARC Inc,WA4TFZ,Albemarle County,38.033340,-78.476520
Alexandria Radio Club,W4HFH,Alexandria County,38.778120,-77.093810
```

### Output

```
~/Projects/ocelotsloth/csmapper/csmapper main*
❯ ./csmapper geojson -u <qrz username> -p <your password> -c clubs_short.csv 
2022/02/21 00:10:46 parsing file
2022/02/21 00:10:46 calling generate geojson
2022/02/21 00:10:46 getting session key
2022/02/21 00:10:47 got session key
2022/02/21 00:10:47 getting session
2022/02/21 00:10:47 got session
2022/02/21 00:10:47 processing clubs
2022/02/21 00:10:47 returning feature set
2022/02/21 00:10:47 marshalling geojson
{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[-78.47555,38.030177]},"properties":{"description":"Albemarle County","title":"Albemarle ARC Inc"}},{"type":"Feature","geometry":{"type":"Point","coordinates":[-77.108402,38.803672]},"properties":{"description":"Alexandria County","title":"Alexandria Radio Club"}}]}
```

### Example Mapped

https://w.gamma.markstenglein.com/mediawiki/index.php/VA_Clubs

