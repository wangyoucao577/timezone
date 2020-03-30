
import argparse
from datetime import datetime, timedelta
import pytz

class MyArgs:

    def __init__(self):
        self.__parser = argparse.ArgumentParser(description='Display date and time.')

        # default positional arguments
        self.__parser.add_argument('-u', '--utc', help='Display UTC time', action='store_true')
        self.__parser.add_argument('-tz', '--tz', help='IANA Time Zone database Name, such as "America/New_York".')
        self.__parser.add_argument('-f', '--format', help='Format string representation of the time, e.g., "%%Y-%%m-%%d %%H:%%M:%%S %%z". See more in "https://docs.python.org/3/library/datetime.html#strftime-strptime-behavior". Some predefined formats: iso,ctime.')
        # parse args
        self.__args = self.__parser.parse_args()

    def get_args(self):
        return self.__args


def get_current_time(utc, tzDataBaseName):
    
    t = datetime.now()
    t = t.astimezone()  # append local timezone

    if utc:
        return t.astimezone(pytz.utc)
    if tzDataBaseName:
        target_timezone = pytz.timezone(tzDataBaseName)
        return t.astimezone(target_timezone)

    return t

def time_to_str(t, format):

    if not format:
        return t.strftime('%a %b %d %T %Z %Y') # default, same with shell `date`'s default format

    if format == 'iso':
        return t.isoformat()
    elif format == 'ctime':
        return t.ctime()
    return t.strftime(format)


def main():
    args = MyArgs().get_args()

    t = get_current_time(args.utc, args.tz)
    print(time_to_str(t, args.format))

if __name__ == '__main__':
    main()
