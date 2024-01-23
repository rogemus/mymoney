from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.mixins import LoginRequiredMixin
from django.shortcuts import redirect, render
from django.utils.decorators import method_decorator
from django.views import View, generic

from .models import Budget, Transaction


@method_decorator(login_required, name="dispatch")
class BudgetList(View):
    template_name = "budget/index.html"

    def get(self, request):
        """
        Budgets list for login user
        """
        budgets = Budget.objects.filter(user=request.user)
        return render(request, self.template_name, {"budgets": budgets})


# [TODO] Add Update funtion for budget
class BudgetDetail(LoginRequiredMixin, generic.DetailView):
    model = Budget
    template_name = "budget/budget-detail.html"


# [TODO] Add update function for transaction
class TransactionDetail(LoginRequiredMixin, generic.DetailView):
    model = Transaction
    template_name = "transaction/transaction-detail.html"


@method_decorator(login_required, name="dispatch")
class TransactionAdd(View):
    def post(self, request, budget_id):
        """
        Add new transaction to budget
        """
        amount = request.POST["amount"]
        is_expense = (
            bool(request.POST["is_expense"]) if "is_expense" in request.POST else False
        )
        description = request.POST["desc"]
        budget = Budget.objects.get(pk=budget_id)

        # [TODO] Add Validation
        transaction = Transaction(
            amount=amount,
            is_expense=is_expense,
            description=description,
            budget=budget,
            user=request.user,
        )
        transaction.save()
        messages.success(request, "Transaction added!")
        return redirect("/budget/%s" % budget_id)


@method_decorator(login_required, name="dispatch")
class BudgetAdd(View):
    template_name = "budget/budget-add.html"

    def get(self, request):
        """
        Render BudgetAdd form
        """
        return render(request, self.template_name)

    def post(self, request):
        """
        Create new budget for login user
        """
        name = request.POST["budget_name"]

        # Add proper validation
        if name == "":
            messages.error(request, "Budget name empty")
        else:
            messages.success(request, "Budget created!")
            new_budget = Budget(name=name, user=request.user)
            new_budget.save()

        return render(request, self.template_name)
