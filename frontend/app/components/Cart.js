let cart = { products: [] };

// Subscriptions
let subs = [];

async function updateCart(data) {
  if (data) {
    cart = data;
    cart.total = cart.products.reduce((acc, pr) => acc + pr.price * pr.count, 0);
    // cart.products = data.products;
    localStorage.setItem('cart', JSON.stringify(cart));
    return subs.forEach((listener) => listener(cart));
  }

  const cartState = JSON.parse(localStorage.getItem('cart'));
  cart = cartState ? cartState : { products: [] };
  cart.total = cart.products.reduce((acc, pr) => acc + pr.price * pr.count, 0);
}

async function getCart() {
  await updateCart();
  return cart;
}

async function setProduct(product, quantity = 1) {
  const isInCart = cart.products.find((pr) => pr.id === product.id);
  if (isInCart) {
    isInCart.count = quantity;
    return updateCart(cart);
  }

  await addProduct(product);
  return setProduct(product, quantity);
}

async function addProduct(product) {
  const products = cart.products;
  const isInCart = products.find((pr) => pr.id === product.id);
  if (isInCart)
    return setProduct(product, isInCart.count + 1);

  cart.products.push({...product, count: 1});
  updateCart(cart);
}

async function purgeFromCart(product) {
  updateCart({ products: cart.products.filter((pr) => pr.id !== product.id) });
}

async function placeOrder({ address, payment }) {
  console.warn('TO BE IMPLEMENTED');
  return false;
}

function subscribe(listener) {
  subs.push(listener);
}

const Cart = {
  update: updateCart,
  get: getCart,
  add: setProduct,
  addMore: addProduct,
  remove: purgeFromCart,
  subscribe: subscribe,
  checkout: placeOrder
};
export default Cart;