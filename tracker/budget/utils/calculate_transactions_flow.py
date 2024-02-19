from ..views.dashboard import BudgetTransation
from .create_empty_transactions_flow import create_empty_transactions_flow
from datetime import datetime


def calculate_transactions_flow(
    transactions: list[BudgetTransation], date: datetime
) -> dict[str, float]:
    """
    Calculate transactions flow for given transactions
    """
    flow = create_empty_transactions_flow(date)

    for t in transactions:
        day = t["transaction__created_at"].day
        month = t["transaction__created_at"].month
        label = f"{day}.{month}"
        flow[label] = round(flow[label] + t["transaction__amount"], 2)

    return flow
