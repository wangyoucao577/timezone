
import argparse
from datetime import datetime, timedelta
import pytz

class MyArgs:

    def __init__(self):
        self.__parser = argparse.ArgumentParser(description='Display date and time.')

        # default positional arguments
        self.__parser.add_argument('-u', '--utc', help='Display UTC time', action='store_true')
        self.__parser.add_argument('-tz', '--tz', help='IANA Time Zone database Name, such as "America/New_York".')

        # parse args
        self.__args = self.__parser.parse_args()

    def get_args(self):
        return self.__args


def get_current_time(utc, tzDataBaseName):
    
    t = datetime.now().astimezone()

    if utc:
        return t.astimezone(pytz.utc)
    if tzDataBaseName:
        target_timezone = pytz.timezone(tzDataBaseName)
        return t.astimezone(target_timezone)

    return t


def main():
    args = MyArgs().get_args()

    t = get_current_time(args.utc, args.tz)
    print(t)

if __name__ == '__main__':
    main()
