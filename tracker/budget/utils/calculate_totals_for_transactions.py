from ..views.types import BudgetTransation


def calculate_totals_for_transactions(transactions: list[BudgetTransation]) -> tuple[float, float, float]:
    """
    Calculate total_income, total_expenses, total (overall) for transactions
    """
    total_expenses = total_income = 0

    for t in transactions:
        if t["transaction__amount"] >= 0:
            total_income += t["transaction__amount"]
        else:
            total_expenses += t["transaction__amount"]

    return total_income, total_expenses, round(total_income + total_expenses, 2)
