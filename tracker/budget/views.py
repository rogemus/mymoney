from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.mixins import LoginRequiredMixin
from django.shortcuts import render
from django.views import generic

from .models import Budget, Transaction


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


@login_required
def add_budget_view(request):
    if request.method == "POST":
        name = request.POST["budget_name"]

        # Add proper validation
        if name == "":
            messages.error(request, "Budget name empty")
        else:
            messages.success(request, "Budget created!")
            new_budget = Budget(name=name)
            new_budget.save()

    return render(request, template_name="budget/budget-add.html")
