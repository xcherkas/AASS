import Home from './pages/Home.svelte'
import Cart from './pages/Cart.svelte'
import Category from './pages/Category.svelte'
import NotFound from './pages/NotFound.svelte'
import Order from './pages/Order.svelte'
import Product from './pages/Product.svelte'
import Profile from './pages/Profile.svelte'
import SignInUp from './pages/SignInUp.svelte'

const routes = {
    // Exact path
    '/': Home,
    '/cart': Cart,
    '/shop/:category?': Category,
    '/order': Order,
    '/product/:id': Product,
    '/profile': Profile,
    '/sign': SignInUp,
    '*': Home
}
export default routes;