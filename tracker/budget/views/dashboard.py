from django.contrib.auth.decorators import login_required
from django.db.models import Prefetch, Q
from django.shortcuts import render
from django.utils import timezone
from django.utils.decorators import method_decorator
from django.views import View

from budget.models import Budget, Transaction
from ..utils import calculate_transactions_flow, group_transactions_by_category, calculate_totals_for_transactions

from .types import BudgetTransation
from typing import cast


@method_decorator(login_required, name="dispatch")
class Dashboard(View):
    template_name = "dashboard/overall.html"

    def get(self, request):
        today = timezone.now()

        transactions = cast(
            list[BudgetTransation],
            list(
                Budget.objects.filter(Q(user=request.user) | Q(shared_to=request.user))
                .prefetch_related(
                    Prefetch(
                        "transaction_set",
                        queryset=Transaction.objects.filter(
                            created_at__year=str(today.year),
                            created_at__month=str(today.month),
                        ).order_by("-created_at"),
                    )
                )
                .values(
                    "transaction",
                    "transaction__amount",
                    "transaction__description",
                    "transaction__created_at",
                    "transaction__unique_id",
                    "transaction__user",
                    "transaction__user__id",
                    "transaction__user__username",
                    "transaction__category",
                    "transaction__category__color",
                    "transaction__category__description",
                    "transaction__category__icon",
                    "transaction__category__name",
                    "transaction__category__unique_id",
                )
                .order_by("transaction__created_at")
            ),
        )

        grouped_by = group_transactions_by_category(transactions)
        transactions_flow = calculate_transactions_flow(transactions, today)
        total_income, total_expenses, total = calculate_totals_for_transactions(transactions)
        context = {
            "total": total,
            "total_income": total_income,
            "total_expenses": total_expenses,
            "latest_transactions": transactions[:10],
            "transactions_per_category": grouped_by,
            "transactions_flow": transactions_flow,
        }
        return render(request, self.template_name, context=context)
