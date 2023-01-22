import FrontPage from './components/pages/FrontPage.vue'
import UtilityBillsPage from './components/pages/UtilityBillsPage.vue'

const routes = [
  {
    name: 'front',
    path: '/',
    component: FrontPage,
    meta: {
      title: 'Главная',
    }
  },
  {
    name: 'utility_bills',
    path: '/utility-bills',
    component: UtilityBillsPage,
    meta: {
      title: 'Коммунальные платежи',
    }
  },
]

export default routes
