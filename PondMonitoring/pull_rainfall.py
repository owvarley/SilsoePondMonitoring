import requests
import json
import os
import time

PATH_DATA = str(os.environ['PATH_DATA'])
DATA_FILE = "rainfall_data"
DATA_FILE_TYPE = ".csv"
RAINFALL_API_CALL = "https://environment.data.gov.uk/flood-monitoring/id/stations/E24712/readings.json?today&_sorted&parameter=rainfall"

class RainfallReading:
    def __init__ (self, time, reading):
        self.time = time
        self.reading = reading

    def __str__ (self):
        return self.time + ", " + str(self.reading)

response = requests.get(RAINFALL_API_CALL)

readings = []

for item in response.json()["items"]:
    readings.append(RainfallReading(item["dateTime"], item["value"]))

filename = PATH_DATA + DATA_FILE + time.strftime("_%Y%m%d") + DATA_FILE_TYPE

with open(filename, "w") as f:
    for reading in readings:
        f.write(str(reading) + "\n")

print ("Total readings: " + str(len(readings)))