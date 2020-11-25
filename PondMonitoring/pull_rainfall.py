import requests
import os
import datetime

PATH_DATA = str(os.environ['PATH_DATA'])
DATA_FILE = "rainfall_data"
DATA_FILE_TYPE = ".csv"
RAINFALL_API_ENDPOINT = "https://environment.data.gov.uk/flood-monitoring/id/stations/E24712/readings.csv?parameter=rainfall&_sorted&date="
ONE_DAY = datetime.timedelta(days=1)

class Rainfall:
    @staticmethod
    def Download():
        now = datetime.datetime.now()
        yesterday = now - ONE_DAY

        yesterday_str = yesterday.strftime("%Y-%m-%d")
        filename_str = ''.join([PATH_DATA, DATA_FILE, "_", yesterday.strftime("%Y%m%d_%H%M%S"), DATA_FILE_TYPE])

        response = requests.get(RAINFALL_API_ENDPOINT + yesterday_str)

        with open(filename_str, "w") as f:
            f.write(response.text)

        print ("Total bytes written: " + str(len(response.content)))