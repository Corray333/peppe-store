import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ShopView from '../views/ShopView.vue'
import AboutView from '../views/AboutView.vue'
import ProductView from '../views/ProductView.vue'
import CartView from '../views/CartView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/products',
      name: 'products',
      component: ShopView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/products/:id',
      name: 'product',
      component: ProductView
    },
    {
      path: '/cart',
      name: 'cart',
      component: CartView
    },
  ]
})

export default router
