from django.contrib import messages
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib.auth.forms import AuthenticationForm, UserCreationForm
from django.contrib.auth.models import User
from django.shortcuts import redirect, render


def signin(request):
    # Add proper validation
    if request.method == "POST":
        password = request.POST["password"]
        username = request.POST["username"]
        user = authenticate(username=username, password=password)

        if user is None:
            messages.error(
                request, "User does not exist. Provide correct username, password"
            )
        else:
            login(request, user)
            return redirect("/budget/list")

    form = AuthenticationForm()
    return render(request, template_name="accounts/login.html", context={"form": form})


def signup(request):
    # Add proper validation
    if request.method == "POST":
        email = request.POST["email"]
        password = request.POST["password"]
        username = request.POST["username"]

        user = User.objects.create_user(username, email, password)

        user.save()
        messages.success(request, "User created!")

    form = UserCreationForm()
    return render(
        request, template_name="accounts/register.html", context={"form": form}
    )


@login_required
def signout(request):
    logout(request)
    redirect("index")
