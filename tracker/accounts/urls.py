from django.urls import path

from . import views

app_name = "accounts"
urlpatterns = [
    path("", views.SignIn.as_view(), name="index"),
    path("signup", views.SignUp.as_view(), name="signup"),
    path(
        "signin",
        views.SignIn.as_view(),
        name="signin",
    ),
    path(
        "signout",
        views.SignOut.as_view(),
        name="signout",
    ),
]
