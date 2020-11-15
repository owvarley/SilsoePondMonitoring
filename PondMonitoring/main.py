import pull_rainfall
import time
import datetime

HOUR_TO_CAPTURE = 16
SLEEP_1_HOUR = 1 * 60 * 60
SLEEP_24_HOUR = 24 * 60 * 60

def main(curr_hour, hour_to_capture, sleep_24, sleep_1):
    while True:
        if curr_hour == hour_to_capture:
            # Pull and save the rainfall data once a day after a given hour to ensure
            # a full set of data is retained
            pull_rainfall.Rainfall.Download()
            time.sleep(sleep_24)
        else:
            time.sleep(sleep_1)

if __name__ == "__main__":
    now = datetime.datetime.now()
    main(now.hour, HOUR_TO_CAPTURE, SLEEP_24_HOUR, SLEEP_1_HOUR)