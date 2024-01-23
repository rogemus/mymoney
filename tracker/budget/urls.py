from django.urls import path

from . import views

app_name = "budget"
urlpatterns = [
    path("list", views.BudgetList.as_view(), name="index"),
    path("add/", views.BudgetAdd.as_view(), name="budget_add"),
    path("<int:pk>/", views.BudgetDetail.as_view(), name="budget_detail"),
    path(
        "t/<int:pk>/",
        views.TransactionDetail.as_view(),
        name="transaction_detail",
    ),
]
