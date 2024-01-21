from django.views import generic

from .models import Budget, Transaction
from django.contrib.auth.mixins import LoginRequiredMixin


class IndexView(LoginRequiredMixin, generic.ListView):
    template_name = "budget/index.html"
    context_object_name = "budgets"

    def get_queryset(self):
        """Return all budgets"""
        return Budget.objects.all()


class BudgetDetail(LoginRequiredMixin, generic.DetailView):
    model = Budget
    template_name = "budget/budget-detail.html"


class TransactionDetail(LoginRequiredMixin, generic.DetailView):
    model = Transaction
    template_name = "transaction/transaction-detail.html"
