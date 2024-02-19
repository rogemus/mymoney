import calendar
from datetime import datetime
from typing import TypedDict

type Transaction = dict[str, str]

class BudgetTransation(TypedDict):
    """
    Class representing Transaction obcject returned from Budget query with .values()
    """

    transaction: str
    transaction__amount: float
    transaction__description: str
    transaction__created_at: datetime
    transaction__unique_id: str
    transaction__user: str
    transaction__user__id: str
    transaction__user__username: str
    transaction__category: str
    transaction__category__color: str
    transaction__category__description: str
    transaction__category__icon: str
    transaction__category__name: str
    transaction__category__unique_id: str


class TransactionGroup(TypedDict):
    """
    Class representing single group of BudgetTransations
    """
    color: str
    description: str
    icon: str
    name: str
    total: float
    unique_id: str
    transactions: list[BudgetTransation]


def group_transactions_by_category(transactions: list[BudgetTransation]) -> dict[str, TransactionGroup]:
    """
    Group transactions based on the `transaction__category__unique_id`. While grouping canculate total amount for group
    """

    grouped_by: dict[str, TransactionGroup] = {}

    for t in transactions:
        cat_id = t["transaction__category__unique_id"]

        if cat_id not in grouped_by:
            grouped_by[cat_id] = {
                "name" : t["transaction__category__name"],
                "description": t["transaction__category__description"],
                "icon": t["transaction__category__icon"],
                "total": t["transaction__amount"],
                "unique_id": t["transaction__category__unique_id"],
                "color": t["transaction__category__color"],
                "transactions": [t]
            }
        else: 
            group = grouped_by[cat_id]
            group['transactions'].append(t)
            amount = t["transaction__amount"]
            total = group["total"]
            group["total"] = round(total + amount, 2)
            grouped_by[cat_id] = group


    return grouped_by


def create_empty_transactions_flow(date: datetime) -> dict[str, float]:
    """
    Create empty data structure that will hold transactions flow in given month
    """
    monthrange = calendar.monthrange(date.year, date.month)
    list_of_days = list(range(1, monthrange[1] + 1))
    list_of_labels = [f"{day}.{date.month}" for day in list_of_days]
    return dict.fromkeys(list_of_labels, 0)


def calculate_transactions_flow(transactions: list[BudgetTransation], date: datetime) -> dict[str, float]:
    """
    Calculate transactions flow for given transactions
    """
    flow = create_empty_transactions_flow(date)

    for t in transactions: 
        day = t["transaction__created_at"].day
        month = t["transaction__created_at"].month
        label = f"{day}.{month}"
        flow[label] = round(
            flow[label] + t["transaction__amount"], 2
        )

    return flow





