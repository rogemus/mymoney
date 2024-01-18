import datetime

from budget.models import Budget, Transaction
from django.test import TestCase
from django.utils import timezone


def create_sample_budget(name):
    return Budget.objects.create(name=name)


class TransactionModel(TestCase):
    def test_transaction_was_made_in_current_month(self):
        """
        in_current_month() return False if transaction was not mode in current month
        """
        test_budget = create_sample_budget("Test Budget")
        transaction_in_past = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now() + datetime.timedelta(days=-90),
        )
        transaction_in_future = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now() + datetime.timedelta(days=90),
        )
        transaction_current_month = Transaction(
            budget=test_budget,
            amount=10,
            description="",
            created_at=timezone.now(),
        )

        self.assertIs(transaction_in_past.in_current_month(), False)
        self.assertIs(transaction_in_future.in_current_month(), False)
        self.assertIs(transaction_current_month.in_current_month(), True)
