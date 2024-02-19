import datetime

from budget.models import Budget, Transaction
from django.contrib.auth.models import User
from django.test import TestCase
from django.utils import timezone


def create_sample_budget(name, user):
    return Budget.objects.create(name=name, description="", user=user)


def create_sample_user(username="user", password="pass"):
    return User.objects.create(username=username, password=password)


# class BudgetListQueary(TestCase):
#     def test_transaction_in_budget_from_current_month(self):
#         """
#
#         """
#         test_user = create_sample_user()
#         test_budget = create_sample_budget("Test Budget", user=test_user)
#         Transaction(
#             budget=test_budget,
#             amount=10,
#             description="",
#             created_at=timezone.now() + datetime.timedelta(days=-90),
#             user=test_user
#         )
#         Transaction(
#             budget=test_budget,
#             amount=10,
#             description="",
#             created_at=timezone.now() + datetime.timedelta(days=90),
#             user=test_user
#         )
#         Transaction(
#             budget=test_budget,
#             amount=10,
#             description="",
#             created_at=timezone.now(),
#             user=test_user
#         )
#         transactions = test_budget.current_month_transaction()
#         print(transactions)


class TransactionModel(TestCase):
    def test_transaction_was_made_in_current_month(self):
        """
        in_current_month() return False if transaction was not mode in current month
        """
        test_user = create_sample_user()
        test_budget = create_sample_budget("Test Budget", user=test_user)
        transaction_in_past = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now() + datetime.timedelta(days=-90),
            user=test_user,
        )
        transaction_in_future = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now() + datetime.timedelta(days=90),
            user=test_user,
        )
        transaction_current_month = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now(),
            user=test_user,
        )

        self.assertIs(transaction_in_past.in_current_month(), False)
        self.assertIs(transaction_in_future.in_current_month(), False)
        self.assertIs(transaction_current_month.in_current_month(), True)
