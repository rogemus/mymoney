from django.contrib import messages
from django.contrib.auth import authenticate, login, logout
from django.contrib.auth.decorators import login_required
from django.contrib.auth.forms import AuthenticationForm, UserCreationForm
from django.contrib.auth.models import User
from django.shortcuts import redirect, render
from django.utils.decorators import method_decorator
from django.views import View


class SignIn(View):
    template_name = "accounts/login.html"
    form = AuthenticationForm()

    def get(self, request):
        return render(request, self.template_name, context={"form": self.form})

    def post(self, request):
        password = request.POST["password"]
        username = request.POST["username"]
        user = authenticate(username=username, password=password)

        if user is not None:
            login(request, user)
            return redirect("/budget/list")
        else:
            messages.error(request, "User does not exst.")

        return render(request, self.template_name, context={"form": self.form})


class SignUp(View):
    template_name = "accounts/register.html"
    form = UserCreationForm()

    def get(self, request):
        return render(request, self.template_name, context={"form": self.form})

    def post(self, request):
        email = request.POST["email"]
        password = request.POST["password"]
        username = request.POST["username"]

        user = User.objects.create_user(username, email, password)

        user.save()
        messages.success(request, "User created!")
        return render(request, self.template_name, context={"form": self.form})


@method_decorator(login_required, name="dispatch")
class SignOut(View):
    def get(self, request):
        logout(request)
        messages.success(request, "User logout!")
        return redirect("/")
