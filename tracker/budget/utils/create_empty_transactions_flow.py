import calendar
from datetime import datetime


def create_empty_transactions_flow(date: datetime) -> dict[str, float]:
    """
    Create empty data structure that will hold transactions flow in given month
    """
    monthrange = calendar.monthrange(date.year, date.month)
    list_of_days = list(range(1, monthrange[1] + 1))
    list_of_labels = [f"{day}.{date.month}" for day in list_of_days]
    return dict.fromkeys(list_of_labels, 0)
