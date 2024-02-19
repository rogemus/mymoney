from typing import TypedDict
from datetime import datetime


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
