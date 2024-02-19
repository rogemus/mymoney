from django.urls import path

from .views import Dashboard

app_name = "budget"
urlpatterns = [
    path(
        "",
       Dashboard.as_view(),
        name="index"
    ),

    path(
        "list/",
        Dashboard.as_view(),
        name="budget_list"
    ),
    path(
        "add/",
        Dashboard.as_view(),
        name="budget_add"
    ),
    path(
        "<int:pk>/",
        Dashboard.as_view(),
        name="budget_detail"
    ),
    path(
        "<int:budget_id>/transaction_add",
        Dashboard.as_view(),
        name="transaction_add",
    ),
    path(
        "<int:budget_id>/share",
        Dashboard.as_view(),
        name="budget_share"
    ),
    path(
        "join/",
        Dashboard.as_view(),
        name="budget_share_token"
    ),
    path(
        "t/<int:pk>/",
        Dashboard.as_view(),
        name="transaction_detail",
    ),
]
