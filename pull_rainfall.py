import requests

response = requests.get("https://environment.data.gov.uk/flood-monitoring/id/stations/E24712/readings.json?today&_sorted&parameter=rainfall")

print (response.status_code)