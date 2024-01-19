from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib.auth.forms import AuthenticationForm
from django.http import HttpResponse
from django.shortcuts import redirect, render


def signin(request):
    if request.user.is_authenticated:
        return redirect("index")

    form = AuthenticationForm()
    return render(request, template_name="accounts/login.html", context={"form": form})


def signup(request):
    return HttpResponse("register")


@login_required(login_url="/accounts/signin")
def signout(request):
    logout(request)
    redirect("index")
