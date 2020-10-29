import requests
import json

response = requests.get("https://environment.data.gov.uk/flood-monitoring/id/stations/E24712/readings.json?today&_sorted&parameter=rainfall")

for item in response.json()["items"]:
    print (item["dateTime"] + " " + str(item["value"]))