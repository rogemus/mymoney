from django.urls import path

from . import views

app_name = "budget"
urlpatterns = [
    path(
        "",
        views.Dashboard.as_view(),
        name="index"
    ),
    path(
        "/list",
        views.BudgetList.as_view(),
        name="budget_list"
    ),
    path(
        "add/",
        views.BudgetAdd.as_view(),
        name="budget_add"
    ),
    path(
        "<int:pk>/",
        views.BudgetDetail.as_view(),
        name="budget_detail"
    ),
    path(
        "<int:budget_id>/transaction_add",
        views.TransactionAdd.as_view(),
        name="transaction_add",
    ),
    path(
        "<int:budget_id>/share",
        views.BudgetShare.as_view(),
        name="budget_share"
    ),
    path(
        "join/",
        views.BudgetShareToken.as_view(),
        name="budget_share_token"
    ),
    path(
        "t/<int:pk>/",
        views.TransactionDetail.as_view(),
        name="transaction_detail",
    ),
]
