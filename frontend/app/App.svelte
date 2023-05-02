<header class="d-flex justify-content-center align-items-center">
  <a href="/" class="header-link" use:link use:active>Home</a>
  <!-- <a href="/hello/user" use:link use:active={'/hello/*'}>Say hi with a default className!</a> -->
  <!-- <a href="/hello/user" use:link use:active>Say hi with all default options!</a> -->
  <!-- <a href="#" class="header-link active">Home</a> -->
  <a href="/shop" class="header-link" use:link use:active={'/shop/*'}>Shop</a>
  <a href="/cart" class="header-link" use:link use:active>Cart ({totalCount})</a>
  {#if logged }
    <a href="/profile" class="header-link" use:link use:active>Profile</a>
  {:else}
    <a href="/sign" class="header-link" use:link use:active>Sign In</a>
  {/if}
</header>

<Router { routes }/>

<footer class="row align-items-stretch">
  <div class="col-12 info">
    <div class="row">
      <div class="col-8 about d-flex align-items-center">
        <p>
          <b>C&T Company</b><br>
          Our company was created about 18 hours ago, but now we reached the TOP-1 position in Tea and Coffee quality competition. We'll deliver all types of beans you want within 24 hours. You are the best, we are too!
        </p>
      </div>
      <div class="col-4 workhours d-flex align-items-center justify-content-center">
        <p>Working hours<br>
        <b>0-24 NONSTOP</b></p>
      </div>
    </div>
  </div>
  <div class="col-12 img"></div>
</footer>

<script>
  import Router from 'svelte-spa-router'
  import { link } from 'svelte-spa-router'
  import active from 'svelte-spa-router/active'
  import routes from './routes'
  import Auth from './components/Auth'
  import Cart from './components/Cart';

  let logged = false;
  let totalCount = 0;

  function calcCartCount(cart) {
    totalCount = cart.products.reduce((acc, item) => acc + item.count, 0);
  }

  Cart.get().then((cart) => {
    calcCartCount(cart);
  });

  Auth.subscribe(() => logged = Auth.isLogged());
  Cart.subscribe((cart) => {
    calcCartCount(cart);
  });


</script>