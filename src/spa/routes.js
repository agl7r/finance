import FrontPage from './components/pages/FrontPage.vue'
import UtilityBillsPage from './components/pages/UtilityBillsPage.vue'

const routes = [
  { name: 'front', path: '/', component: FrontPage },
  { name: 'utility_bills', path: '/utility-bills', component: UtilityBillsPage },
]

export default routes
