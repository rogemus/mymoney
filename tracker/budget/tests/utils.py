from datetime import datetime
from ..utils import (
    calculate_transactions_flow,
    create_empty_transactions_flow,
    group_transactions_by_category,
    calculate_totals_for_transactions,
)
from ..views.types import BudgetTransation
from django.test import TestCase


def create_mock_transaction(amount, id, date) -> BudgetTransation:
    return {
        "transaction": "",
        "transaction__amount": amount,
        "transaction__description": "",
        "transaction__created_at": date,
        "transaction__unique_id": "",
        "transaction__user": "",
        "transaction__user__id": "",
        "transaction__user__username": "",
        "transaction__category": "",
        "transaction__category__color": "",
        "transaction__category__description": "",
        "transaction__category__icon": "",
        "transaction__category__name": "",
        "transaction__category__unique_id": id,
    }


class Utils(TestCase):
    def test_calculate_totals_for_transactions(self):
        """
        calculate_totals_for_transactions() returns total_income (positive), total_expences (negative), total (overall) for transactions
        """

        t_1 = create_mock_transaction(12, "cat_id_1", datetime(2024, 2, 1))
        t_2 = create_mock_transaction(-3.2, "cat_id_1", datetime(2024, 2, 2))
        t_3 = create_mock_transaction(5, "cat_id_1", datetime(2024, 2, 3))
        t_4 = create_mock_transaction(-2.16, "cat_id_1", datetime(2024, 2, 3))
        transactions = [t_1, t_2, t_3, t_4]
        total_income, total_expences, total = calculate_totals_for_transactions(transactions)

        self.assertEqual(total_income, 17)
        self.assertEqual(total_expences, -5.36)
        self.assertEqual(total, 11.64)

    def test_calculate_transactions_flow(self):
        """
        calculate_transactions_flow() return dict for each day for given date with calculated total amount for each day
        """

        t_1 = create_mock_transaction(12, "cat_id_1", datetime(2024, 2, 1))
        t_2 = create_mock_transaction(3, "cat_id_1", datetime(2024, 2, 2))
        t_3_1 = create_mock_transaction(5, "cat_id_1", datetime(2024, 2, 3))
        t_3_2 = create_mock_transaction(3.25, "cat_id_1", datetime(2024, 2, 3))

        transactions = [t_1, t_2, t_3_1, t_3_2]
        flow = calculate_transactions_flow(transactions, datetime(2024, 2, 23))
        self.assertDictEqual(
            flow,
            {
                "1.2": 12,
                "10.2": 0,
                "11.2": 0,
                "12.2": 0,
                "13.2": 0,
                "14.2": 0,
                "15.2": 0,
                "16.2": 0,
                "17.2": 0,
                "18.2": 0,
                "19.2": 0,
                "2.2": 3,
                "20.2": 0,
                "21.2": 0,
                "22.2": 0,
                "23.2": 0,
                "24.2": 0,
                "25.2": 0,
                "26.2": 0,
                "27.2": 0,
                "28.2": 0,
                "29.2": 0,
                "3.2": 8.25,
                "4.2": 0,
                "5.2": 0,
                "6.2": 0,
                "7.2": 0,
                "8.2": 0,
                "9.2": 0,
            },
        )

    def test_group_transactions_by_category(self):
        """
        group_transaction_by_category() returns dict of category with transactions.
        While grouping total for the category is calculated
        """

        tc1_1 = create_mock_transaction(12, "cat_id_1", datetime.now())
        tc1_2 = create_mock_transaction(3, "cat_id_1", datetime.now())
        tc2_1 = create_mock_transaction(5, "cat_id_2", datetime.now())

        transactions = [tc1_1, tc1_2, tc2_1]
        grouped = group_transactions_by_category(transactions)
        self.assertIs(grouped["cat_id_1"]["unique_id"], "cat_id_1")
        self.assertIs(grouped["cat_id_1"]["total"], 15)
        self.assertIs(len(grouped["cat_id_1"]["transactions"]), 2)

        self.assertIs(grouped["cat_id_2"]["unique_id"], "cat_id_2")
        self.assertIs(grouped["cat_id_2"]["total"], 5)
        self.assertIs(len(grouped["cat_id_2"]["transactions"]), 1)

    def test_create_empty_transactions_flow(self):
        """
        create_empty_transactions_flow() return dict for each day for given date
        """
        feb_29 = datetime(2024, 2, 1)
        feb_28 = datetime(2023, 2, 1)
        jan = datetime(2024, 1, 1)
        apr = datetime(2024, 4, 1)

        feb_29_flow = create_empty_transactions_flow(feb_29)
        feb_28_flow = create_empty_transactions_flow(feb_28)
        jan_flow = create_empty_transactions_flow(jan)
        apr_flow = create_empty_transactions_flow(apr)
        self.assertDictEqual(
            feb_29_flow,
            {
                "1.2": 0,
                "2.2": 0,
                "3.2": 0,
                "4.2": 0,
                "5.2": 0,
                "6.2": 0,
                "7.2": 0,
                "8.2": 0,
                "9.2": 0,
                "10.2": 0,
                "11.2": 0,
                "12.2": 0,
                "13.2": 0,
                "14.2": 0,
                "15.2": 0,
                "16.2": 0,
                "17.2": 0,
                "18.2": 0,
                "19.2": 0,
                "20.2": 0,
                "21.2": 0,
                "22.2": 0,
                "23.2": 0,
                "24.2": 0,
                "25.2": 0,
                "26.2": 0,
                "27.2": 0,
                "28.2": 0,
                "29.2": 0,
            },
        )
        self.assertDictEqual(
            feb_28_flow,
            {
                "1.2": 0,
                "2.2": 0,
                "3.2": 0,
                "4.2": 0,
                "5.2": 0,
                "6.2": 0,
                "7.2": 0,
                "8.2": 0,
                "9.2": 0,
                "10.2": 0,
                "11.2": 0,
                "12.2": 0,
                "13.2": 0,
                "14.2": 0,
                "15.2": 0,
                "16.2": 0,
                "17.2": 0,
                "18.2": 0,
                "19.2": 0,
                "20.2": 0,
                "21.2": 0,
                "22.2": 0,
                "23.2": 0,
                "24.2": 0,
                "25.2": 0,
                "26.2": 0,
                "27.2": 0,
                "28.2": 0,
            },
        )
        self.assertDictEqual(
            jan_flow,
            {
                "1.1": 0,
                "2.1": 0,
                "3.1": 0,
                "4.1": 0,
                "5.1": 0,
                "6.1": 0,
                "7.1": 0,
                "8.1": 0,
                "9.1": 0,
                "10.1": 0,
                "11.1": 0,
                "12.1": 0,
                "13.1": 0,
                "14.1": 0,
                "15.1": 0,
                "16.1": 0,
                "17.1": 0,
                "18.1": 0,
                "19.1": 0,
                "20.1": 0,
                "21.1": 0,
                "22.1": 0,
                "23.1": 0,
                "24.1": 0,
                "25.1": 0,
                "26.1": 0,
                "27.1": 0,
                "28.1": 0,
                "29.1": 0,
                "30.1": 0,
                "31.1": 0,
            },
        )
        self.assertDictEqual(
            apr_flow,
            {
                "1.4": 0,
                "2.4": 0,
                "3.4": 0,
                "4.4": 0,
                "5.4": 0,
                "6.4": 0,
                "7.4": 0,
                "8.4": 0,
                "9.4": 0,
                "10.4": 0,
                "11.4": 0,
                "12.4": 0,
                "13.4": 0,
                "14.4": 0,
                "15.4": 0,
                "16.4": 0,
                "17.4": 0,
                "18.4": 0,
                "19.4": 0,
                "20.4": 0,
                "21.4": 0,
                "22.4": 0,
                "23.4": 0,
                "24.4": 0,
                "25.4": 0,
                "26.4": 0,
                "27.4": 0,
                "28.4": 0,
                "29.4": 0,
                "30.4": 0,
            },
        )
