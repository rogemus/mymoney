# Generated by Django 5.0.1 on 2024-01-18 21:24

from django.db import migrations


class Migration(migrations.Migration):
    dependencies = [
        ("budget", "0004_rename_is_expance_transaction_is_expanse"),
    ]

    operations = [
        migrations.RenameField(
            model_name="transaction",
            old_name="is_expanse",
            new_name="is_expense",
        ),
    ]