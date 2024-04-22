import { Routes } from '@angular/router';
import { LoginPage } from './routes/login/login.component';
import { BudgetListPage } from './routes/budgetList/budgetList.component';
import { BudgetDetailsPage } from './routes/budgetDetails/budgetDetails.component';

export const routes: Routes = [
  {
    path: '', component: LoginPage,
  },
  {
    path: 'budgets', component: BudgetListPage,
  },
  {
    path: 'budget', component: BudgetDetailsPage
  }
];
