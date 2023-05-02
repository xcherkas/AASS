<div class="container cart-item">
  <div class="row">
    <div class="col-12">
      <div class="row text-center progress-info justify-content-center">
        <div class="col-12 col-md-4"><a href="/cart" use:link class="text-success"><b>Order summary</b></a></div>
        <div class="col-12 col-md-4"><a href="/order" use:link class="text-success">Delivery&Payment</a></div>
        <div class="col-12 col-md-4"><span class="text-success">Confirm checkout</span></div>
        <div class="col-12">
          <div class="progress">
            <div class="progress-bar bg-success" role="progressbar" style="width: 33%" aria-valuenow="33" aria-valuemin="0" aria-valuemax="100"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="row cart-list">
    <div class="col-12">
      <table class="table table-borderless">
        <tbody>
          {#each cart.products as item}
            <tr>
              <th scope="row">
                <img src="{ item.img }" alt="Image of { item.name }">
              </th>
              <td>{ item.title }</td>
              <td>${ item.price }&nbsp;x</td>
              <td>
                <div class="input-group">
                  <div class="input-group-prepend">
                    <button class="btn btn-success" type="button" on:click={ count.bind(null, item, Math.max(1, item.count - 1)) }>-</button>
                  </div>
                  <input type="text" class="form-control" readonly placeholder="Count" bind:value={ item.count }>
                  <div class="input-group-append">
                    <button class="btn btn-success" type="button" on:click={ count.bind(null, item, item.count + 1) }>+</button>
                  </div>
                </div>
              </td>
              <td>${ item.price * item.count }</td>
              <td on:click={ ()=>{ Cart.remove(item) } }>&times;</td>
            </tr>
          {:else}
            <tr>
              <td colspan="6" class="text-center">No products here :(</td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
    <div class="col-12 d-flex justify-content-end align-items-center">
      <span class="total">Total: ${ cart.total || 0 }</span>
      {#if cart.products.length > 0}
        <button type="button" class="btn btn-success" on:click={ order }>Checkout</button>
      {/if}
    </div>
  </div>
</div>

<script>
  import axios from 'axios';
  import of from 'await-of';
  import Cart from '../components/Cart';
  import { push, link } from 'svelte-spa-router'
  import { onMount } from 'svelte';

  let cart = {
    products: { data: [] }
  };

  function order() {
    console.log('New Order', cart);
    Cart.update({ products: [] });
    push('/');
  }

  async function count(item, count) {
    await Cart.add(item, count);
  }
  onMount(async () => {
    cart = await Cart.get();
    Cart.subscribe(async (cartUpdate) => { cart = cartUpdate });
  });
</script>