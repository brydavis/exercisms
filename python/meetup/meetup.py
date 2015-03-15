from datetime import date, timedelta

MeetupDayException = Exception(" error")

weekday_digit = {
    "Monday": 1,
    "Tuesday": 2,
    "Wednesday": 3,
    "Thursday": 4,
    "Friday": 5,
    "Saturday": 6,
    "Sunday": 7,
}

days_per_month = [31,28,31,30,31,30,31,31,30,31,30,31]

def add_gigasecond(date):
	return date + timedelta(seconds=10**9)


def meetup_day(year, month, weekday, week):
    # weekday = weekday_digit[weekday]
    # for day in range(1, days_per_month[month]):
    try:
       date(year, month, day)
    except Exception, e:
        raise e

