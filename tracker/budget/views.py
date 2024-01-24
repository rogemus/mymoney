from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.mixins import LoginRequiredMixin
from django.contrib.auth.models import User
from django.shortcuts import redirect, render
from django.utils.decorators import method_decorator
from django.views import View, generic

from .models import Budget, Transaction


@method_decorator(login_required, name="dispatch")
class BudgetList(View):
    template_name = "budget/budget-list.html"

    def get(self, request):
        """
        Budgets list for login user
        """
        budgets = Budget.objects.filter(user=request.user)
        shared_budgets = Budget.objects.filter(shared_to=request.user)

        return render(
            request,
            self.template_name,
            {"budgets": budgets, "shared_budgets": shared_budgets},
        )


class BudgetDetail(View):
    template_name = "budget/budget-detail.html"

    def get(self, request, pk):
        budget = Budget.objects.get(pk=pk)
        shared_to_users = budget.shared_to.all()

        print(shared_to_users)
        return render(
            request,
            self.template_name,
            context={"budget": budget, "shared_to_users": shared_to_users},
        )


# [TODO] Add update function for transaction
class TransactionDetail(LoginRequiredMixin, generic.DetailView):
    model = Transaction
    template_name = "transaction/transaction-detail.html"


class BudgetShare(View):
    def post(self, request, budget_id):
        # TODO: this should be 2 step feature.
        # User A types user emails in field on front, and User B should accept link in the email
        share_to = request.POST["share_to"]
        budget = Budget.objects.get(pk=budget_id)
        user_share_to = User.objects.get(pk=share_to)
        budget.shared_to.add(user_share_to)
        budget.save()
        messages.success(request, "Budget shared")
        return redirect("/budget/%s" % budget_id)


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
