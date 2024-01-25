from django.http.response import HttpResponse
from django.shortcuts import render
from django.views import View


# TODO: split this into multiple files
class BudgetListRest(View):
    def get(self, request):
        """
        get budget list
        """
        return HttpResponse("Budget list")

    def post(self, request):
        """
        create budget
        """
        return HttpResponse("post Budget Details")


class BudgetDetailRest(View):
    def get(self, request, budget_id):
        """
        get budget
        """
        return HttpResponse("get Budget Details")

    def delete(self, request, budget_id):
        """
        delete budget
        """
        return HttpResponse("get Budget Details")

    def put(self, request, budget_id):
        """
        update budget
        """
        return HttpResponse("get Budget Details")


class BudgetShareRest(View):
    def post(self, request, budget_id):
        """
        share budget
        """
        return HttpResponse("share budget")


class TransactionListRest(View):
    def get(self, request):
        """
        get transaction list
        """
        return HttpResponse("get transaction list")

    def post(self, request):
        """
        create transaction
        """
        return HttpResponse("post transaction")


class TransactionDetailRest(View):
    def get(self, request, transaction_id):
        """
        get transaction
        """
        return HttpResponse("get transaction")

    def delete(self, request, transaction_id):
        """
        delete transaction
        """
        return HttpResponse("delete transaction")

    def put(self, request, transaction_id):
        """
        update transaction
        """
        return HttpResponse("put transaction")


class UserRest(View):
    def post(self, request):
        """
        create user
        """
        return HttpResponse("post user")


class UserDetailRest(View):
    def get(self, request, user_id):
        """
        get user
        """
        return HttpResponse("get user")

    def delete(self, request, user_id):
        """
        delete user
        """
        return HttpResponse("delete user")

    def put(self, request, user_id):
        """
        update user
        """
        return HttpResponse("put user")
