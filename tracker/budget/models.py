import calendar
from datetime import date

from django.contrib import admin
from django.contrib.auth.models import User
from django.db import models


class Budget(models.Model):
    def __str__(self):
        return self.name

    name = models.CharField(max_length=150)
    created_at = models.DateTimeField(auto_now_add=True)
    user = models.ForeignKey(User, on_delete=models.DO_NOTHING)
    description = models.CharField(max_length=250, default="")


class Transaction(models.Model):
    def __str__(self):
        return "%s: %s, %s" % (self.budget.id, self.amount, self.is_expense)

    @admin.display(
        boolean=True,
        description="In Current Month?",
    )
    def in_current_month(self):
        now = date.today()
        month = calendar.monthrange(now.year, now.month)
        start_date = now.replace(day=1)
        end_date = now.replace(day=month[1])

        return start_date <= self.created_at.date() <= end_date

    budget = models.ForeignKey(Budget, on_delete=models.CASCADE)
    created_at = models.DateTimeField(auto_now_add=True)
    amount = models.FloatField(null=True, blank=True, default=0.0)
    is_expense = models.BooleanField(default=True)
    description = models.CharField(max_length=300, default="")
    user = models.ForeignKey(User, on_delete=models.DO_NOTHING)


# class TransactionCategory(models.Model):
#     transaction = models.ForeignKey(Transaction, on_delete=)
#     name = models.CharField(max_length=100)
#     icon = models.CharField(max_length=200)
#     description = models.CharField(max_length=300)
