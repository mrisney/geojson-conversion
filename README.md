# geojson-conversion

## Overview
A simple GeoJSON serializer / deserializer

Use it as a simple explanation of how to cpnvert a .csv file of long/lat cooridnates to a GeoJSON Feature Collection

```
go mod init geojson
go build
./geojson
```
When used from the command line, it attempts to read the input file, write out a file points.geojson

### What does it not do well?
1. It is not designed to validate input. The program assumes that the input is valid .csv Behavior for invalid coordinbates are undefined and could panic.
