from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.mixins import LoginRequiredMixin
from django.shortcuts import render
from django.utils.decorators import method_decorator
from django.views import View, generic

from .models import Budget, Transaction


@method_decorator(login_required, name="dispatch")
class BudgetList(View):
    template_name = "budget/index.html"

    def get(self, request):
        budgets = Budget.objects.filter(user=request.user)
        return render(request, self.template_name, {"budgets": budgets})


class BudgetDetail(LoginRequiredMixin, generic.DetailView):
    model = Budget
    template_name = "budget/budget-detail.html"


class TransactionDetail(LoginRequiredMixin, generic.DetailView):
    model = Transaction
    template_name = "transaction/transaction-detail.html"


@method_decorator(login_required, name="dispatch")
class BudgetAdd(View):
    template_name = "budget/budget-add.html"

    def get(self, request):
        return render(request, self.template_name)

    def post(self, request):
        name = request.POST["budget_name"]

        # Add proper validation
        if name == "":
            messages.error(request, "Budget name empty")
        else:
            messages.success(request, "Budget created!")
            new_budget = Budget(name=name, user=request.user)
            new_budget.save()

        return render(request, self.template_name)
