import axios from 'axios';
import of from 'await-of';
import Cookies from 'js-cookie';

// Token variable
let token = Cookies.get('CT-token');

// Interceptor
axios.interceptors.request.use((config) => {
  if (token)
    config.headers['Authorization'] = `Bearer ${ token }`;
  return config;
});

// Subscriptions
let subs = [];

const Auth = {

  auth: async ({ email, password }) => {
    let [res,] = await of(axios.post('/api/user/login', {
      email: email,
      password: password
    }));
    token = res.data.access_token;
    Cookies.set('CT-token', token, {
      expires: 1/145 // 10 min
    })
    subs.forEach((listener) => listener());
    return true;
  },

  register: async ({ email, name, phone, password, password_confirmation }) => {
    let [res,] = await of(axios.post('/api/user/register', {
      email: email,
      password: password,
      password_confirmation: password_confirmation,
      name: name,
      phone: phone
    }));
    token = res.data.access_token;
    subs.forEach((listener) => listener());
    return true;
  },

  profile: async () => {
    if (!Auth.isLogged())
      return {};

    let [res,] = await of(axios.post('/api/user/get'));
    return res ? res.data.data : {};
  },

  deAuth: async () => {
    token = undefined;
    Cookies.remove('CT-token');
    subs.forEach((listener) => listener());
  },

  isLogged: function () {
    return token != undefined;
  },

  subscribe: function (listener) {
    subs.push(listener);
    if (token)
      listener();
  }

};
export default Auth;