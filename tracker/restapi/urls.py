from django.urls import path

from . import views

app_name = "restapi"
urlpatterns = [
    path(
        "budgets",
        views.BudgetListRest.as_view(),
        name="budget_list"
    ),
    path(
        "budget/<int:budget_id>",
        views.BudgetDetailRest.as_view(),
        name="budget_details",
    ),
    path(
        "budget/<int:budget_id>/share",
        views.BudgetShareRest.as_view(),
        name="budget_share",
    ),
    path(
        "transactions",
        views.TransactionListRest.as_view(),
        name="transaction"
    ),
    path(
        "transaction/<int:transaction_id>",
        views.TransactionDetailRest.as_view(),
        name="transaction_details",
    ),
    path(
        "user",
        views.UserRest.as_view(),
        name="user"
    ),
    path(
        "user/<int:user_id>",
        views.UserDetailRest.as_view(),
        name="user_detail"
    ),
]
