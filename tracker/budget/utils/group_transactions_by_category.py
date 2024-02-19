from ..views.dashboard import BudgetTransation
from typing import TypedDict


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


def group_transactions_by_category(
    transactions: list[BudgetTransation],
) -> dict[str, TransactionGroup]:
    """
    Group transactions based on the `transaction__category__unique_id`. While grouping canculate total amount for group
    """

    grouped_by: dict[str, TransactionGroup] = {}

    for t in transactions:
        cat_id = t["transaction__category__unique_id"]

        if cat_id not in grouped_by:
            grouped_by[cat_id] = {
                "name": t["transaction__category__name"],
                "description": t["transaction__category__description"],
                "icon": t["transaction__category__icon"],
                "total": t["transaction__amount"],
                "unique_id": t["transaction__category__unique_id"],
                "color": t["transaction__category__color"],
                "transactions": [t],
            }
        else:
            group = grouped_by[cat_id]
            group["transactions"].append(t)
            amount = t["transaction__amount"]
            total = group["total"]
            group["total"] = round(total + amount, 2)
            grouped_by[cat_id] = group

    return grouped_by
