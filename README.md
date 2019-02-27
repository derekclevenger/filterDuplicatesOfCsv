# Filters out duplicate urls in a CSV file
## For running locally 
### Argument 1 in the cli is the location of the file to read in
### Argurment 2 in the cli is the name of the file you want to write to
### Execute: go run main.go [your/path/to/csv/file.csv] [your/path/to/desired/end/location.csv] "[Desired url to filter on]"


# For Docker
###  Change to directory containing files
### Run: docker build --tag filterduplicates .
### Run: docker run -t -i -v [path to local FOLDER containing csv]:/var/[folder containing csv name] filterduplicates /bin/bash
### go run main.go /var/[foldername]/[name of csv file] /var/[foldername]/[name of resulting csv file] "[Desired url to filter on]"
### Run: exit

