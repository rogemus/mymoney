from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib.auth.forms import AuthenticationForm, UserCreationForm
from django.contrib.auth.models import User
from django.http import HttpResponse
from django.shortcuts import redirect, render


def signin(request):
    if request.user.is_authenticated:
        return redirect("/budget/")

    form = AuthenticationForm()
    return render(request, template_name="accounts/login.html", context={"form": form})


def register_user(request):
    password = request.POST["password"]
    email = request.POST["email"]
    username = request.POST["username"]

    user = User.objects.create_user(username, email, password)

    user.save()
    return HttpResponse("User created")


def signup(request):
    form = UserCreationForm()
    return render(
        request, template_name="accounts/register.html", context={"form": form}
    )


@login_required(login_url="/accounts/signin")
def signout(request):
    logout(request)
    redirect("index")
