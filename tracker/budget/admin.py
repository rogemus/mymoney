from django.contrib import admin

from .models import Budget, Invitation, Transaction, TransactionCategory


class TransactionAdmin(admin.ModelAdmin):
    list_display = ["amount", "created_at", "in_current_month", "user"]
    list_filter = ["created_at", "user"]


class TransactionInline(admin.TabularInline):
    model = Transaction
    extra = 0


class BudgetAdmin(admin.ModelAdmin):
    fields = ["name", "user", "shared_to"]
    inlines = [TransactionInline]
    search_fields = ["name"]


class InvitationsAdmin(admin.ModelAdmin):
    model = Invitation


class TransactionCategoryAdmin(admin.ModelAdmin):
    model = TransactionCategory


admin.site.register(Invitation, InvitationsAdmin)
admin.site.register(Budget, BudgetAdmin)
admin.site.register(Transaction, TransactionAdmin)
admin.site.register(TransactionCategory, TransactionCategoryAdmin)
