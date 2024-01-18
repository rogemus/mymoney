from django.urls import path

from . import views

app_name = "budget"
urlpatterns = [
    path("", views.IndexView.as_view(), name="index"),
    path("<int:pk>/", views.BudgetDetail.as_view(), name="budget_detail"),
    path(
        "t/<int:pk>/",
        views.TransactionDetail.as_view(),
        name="transaction_detail",
    ),
]
