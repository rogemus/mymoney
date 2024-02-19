from django.urls import path

from .views import Dashboard

app_name = "budget"
urlpatterns = [
    path(
        "",
       Dashboard.as_view(),
        name="index"
    ),
]
