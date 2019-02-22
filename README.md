# Filters out duplicate urls in a CSV file
## For running locally 
### Argument 1 in the cli is the location of the file to read in
#### Argurment 2 in the cli is the name of the file you want to write to
### Execute: go run main.go [your/path/to/csv/file.scv] [your/path/to/desired/end/location.csv]

# For Docker
###  Change to directory container files
### Run: docker build --tag filterduplicates .
### Run: docker run filterDuplicates /bin/bash
### Run: /bin/bash
### Execute: go run main.go [your/path/to/csv/file.scv] [your/path/to/desired/end/location.csv]

