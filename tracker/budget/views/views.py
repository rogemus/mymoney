import calendar

from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.models import User
from django.db.models import Prefetch, Q
from django.shortcuts import redirect, render
from django.utils import timezone
from django.utils.decorators import method_decorator
from django.views import View, generic

from .models import Budget, Invitation, Transaction

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



@method_decorator(login_required, name="dispatch")
class Dashboard(View):
    template_name = "dashboard/overall.html"

    def get(self, request):
        today = timezone.now()

        transactions = list(
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
        )

        for t in transactions:
            if t["transaction"] is not None:
                grouped_transaction_by_catergory(t, grouped_by)
                calculate_transactions_flow(t, transactions_flow)

                if t["transaction__amount"] >= 0:
                    total_income += t["transaction__amount"]
                else:
                    total_expenses += t["transaction__amount"]

        context = {
            "total": round(total_income + total_expenses, 2),
            "total_income": round(total_income, 2),
            "total_expenses": round(total_expenses, 2),
            "latest_transactions": transactions[:10],
            "transactions_per_category": grouped_by,
            "transactions_flow": transactions_flow,
        }
        return render(request, self.template_name, context=context)
