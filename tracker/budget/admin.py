from django.contrib import admin

from .models import Budget, Transaction


class TransactionAdmin(admin.ModelAdmin):
    list_display = ["amount", "is_expense", "created_at", "in_current_month"]
    list_filter = ["created_at"]


class TransactionInline(admin.TabularInline):
    model = Transaction
    extra = 0


class BudgetAdmin(admin.ModelAdmin):
    fields = ["name"]
    inlines = [TransactionInline]
    search_fields = ["name"]


admin.site.register(Budget, BudgetAdmin)
admin.site.register(Transaction, TransactionAdmin)
