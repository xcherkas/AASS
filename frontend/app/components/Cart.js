import axios from 'axios';
import of from 'await-of';

let cart = undefined;

// Subscriptions
let subs = [];

async function updateCart(data) {
  if (data) {
    cart = data;
    cart.products.data = cart.products.data.map(el => {
      let img = '//placehold.it/200';
      if (el.images.data[0])
        img = el.images.data[0].link;

      return {...el, img: img };
    });
    return subs.forEach((listener) => listener(cart));
  }

  let [res, err] = await of(axios.post('/api/order/get-cart/'));
  if (err)
    cart = { products: { data: [] } }
  else {
    cart = res.data.data;
  }

  cart.products.data = cart.products.data.map(el => {
    let img = '//placehold.it/200';
    if (el.images.data[0])
      img = el.images.data[0].link;

    return {...el, img: img };
  });
}

async function getCart() {
  await updateCart();
  return cart;
}

async function setProduct(productId, quantity = 1) {
  let [res,] = await of(axios.post('/api/order/update-count/', {
    product_id: productId,
    count: quantity
  }));
  updateCart(res.data.data);
}

async function addProduct(productId) {
  let [res,] = await of(axios.post('/api/order/add-to-cart/', {
    product_id: productId
  }));
  updateCart(res.data.data);
}

async function purgeFromCart(productId) {
  let [res,] = await of(axios.post('/api/order/delete-from-cart/', {
    products: [productId]
  }));
  updateCart(res.data.data);
}

async function placeOrder({ address, payment }) {
  if (!address || !payment)
    return false;

  let [res, err] = await of(axios.post('/api/order/checkout/', {
    address: JSON.stringify(address),
    paymentData: JSON.stringify(payment)
  }));

  if (err)
    return false;

  updateCart(res.data.data);
  return true;
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