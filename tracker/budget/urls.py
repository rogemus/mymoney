from django.urls import path

from . import views

app_name = "budget"
urlpatterns = [
    path("", views.BudgetList.as_view(), name="index"),
    path("add/", views.BudgetAdd.as_view(), name="budget_add"),
    path("<int:pk>/", views.BudgetDetail.as_view(), name="budget_detail"),
    path(
        "<int:budget_id>/transaction_add",
        view=views.TransactionAdd.as_view(),
        name="transaction_add",
    ),
    path(
        "t/<int:pk>/",
        views.TransactionDetail.as_view(),
        name="transaction_detail",
    ),
]
