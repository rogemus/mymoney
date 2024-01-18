from django.views import generic

from .models import Budget, Transaction


class IndexView(generic.ListView):
    template_name = "budget/index.html"
    context_object_name = "budgets"

    def get_queryset(self):
        """Return all budgets"""
        return Budget.objects.all()


class BudgetDetail(generic.DetailView):
    model = Budget
    template_name = "budget/budget-detail.html"


class TransactionDetail(generic.DetailView):
    model = Transaction
    template_name = "transaction/transaction-detail.html"
