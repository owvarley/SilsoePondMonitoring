import pull_rainfall
import time
import datetime
import os

HOUR_TO_CAPTURE = int(os.environ['HOUR_TO_CAPTURE'])
SLEEP_1_HOUR = 1 * 60 * 60
SLEEP_24_HOUR = 24 * 60 * 60

def main(curr_hour, hour_to_capture, sleep_24, sleep_1):
    while True:
        if curr_hour >= hour_to_capture:
            # Pull and save the rainfall data once a day after a given hour to ensure
            # a full set of data is retained
            print (str(datetime.datetime.now()) + " capturing data.")
            pull_rainfall.Rainfall.Download()
            print ("Waiting...")
            time.sleep(sleep_24)
        else:
            time.sleep(sleep_1)

if __name__ == "__main__":
    print ("Starting Pond Monitor...")
    print ("Hour to Capture: " + str(HOUR_TO_CAPTURE))
    now = datetime.datetime.now()
    main(now.hour, HOUR_TO_CAPTURE, SLEEP_24_HOUR, SLEEP_1_HOUR)