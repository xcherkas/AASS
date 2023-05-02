<div class="container cart-item">
  <div class="row">
    <div class="col-12">
      <div class="row text-center progress-info justify-content-center">
        <div class="col-12 col-md-4"><a href="/cart" use:link class="text-success">Order summary</a></div>

        {#if step == 1}
          <div class="col-12 col-md-4"><span class="text-success"><b>Delivery&Payment</b></span></div>
        {:else}
          <div class="col-12 col-md-4"><span class="text-success">Delivery&Payment</span></div>
        {/if}

        {#if step == 2}
          <div class="col-12 col-md-4"><span class="text-success"><b>Confirm checkout</b></span></div>
        {:else}
          <div class="col-12 col-md-4"><span class="text-success">Confirm checkout</span></div>
        {/if}

        <div class="col-12">
          <div class="progress">
            <div class="progress-bar bg-success" role="progressbar" style="width: 66%" bind:this={progressbar}></div>
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
                <img src="{ item.img }" alt="Image #1">
              </th>
              <td>{ item.name }</td>
              <td>${ item.cost }&nbsp;x</td>
              <td>
                <div class="input-group">
                  <input type="text" readonly class="form-control" placeholder="Count" bind:value={ item.count }>
                </div>
              </td>
              <td>${ item.cost * item.count }</td>
              <td></td>
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
      <span class="total">Total: { cart.total || 0 }</span>
    </div>
  </div>

  <form class="row delivery-list" on:submit|preventDefault={ placeOrder }>
    <div class="col-12 col-md-6 userinfo">
      <h3>User information</h3>
      <div class="form-group row">
        <label class="col-sm-4 col-form-label">Full name</label>
        <div class="col-sm-8">
          <input type="text" class="form-control" placeholder="Full name" readonly bind:value={ profile.name }>
        </div>
      </div>
      <div class="form-group row">
        <label class="col-sm-4 col-form-label">Email</label>
        <div class="col-sm-8">
          <input type="email" class="form-control" placeholder="Email" readonly bind:value={ profile.email }>
        </div>
      </div>
      <div class="form-group row">
        <label class="col-sm-4 col-form-label">Phone number</label>
        <div class="col-sm-8">
          <input type="tel" class="form-control" placeholder="Phone number" readonly bind:value={ profile.phone }>
        </div>
      </div>
      <div class="form-group row">
        <label class="col-sm-4 col-form-label">Address</label>
        <div class="col-sm-8">
          <input type="text" class="form-control" placeholder="Address" readonly={step == 2} bind:value={ profile.address }>
        </div>
      </div>
    </div>
    <div class="col-12 col-md-6 payment">
      <h3>Payment method</h3>
      <div class="form-group row">
        <label class="col-sm-4 col-form-label">Choose one</label>
        <div class="col-sm-8">
          <div class="form-check">
            <input class="form-check-input" type="radio" name="payment" value="IBAN" disabled={step == 2} bind:group={ payment.method }>
            <label class="form-check-label">
              Pay with SEPA Direct Debit
            </label>
          </div>
          <div class="form-check">
            <input class="form-check-input" type="radio" name="payment" value="CARD" disabled={step == 2} bind:group={ payment.method }>
            <label class="form-check-label">
              Pay by credit card
            </label>
          </div>
        </div>
      </div>
      {#if payment.method == 'IBAN' }
        <div class="form-group row">
          <label class="col-sm-4 col-form-label">Your IBAN</label>
          <div class="col-sm-8">
            <input type="text" class="form-control" placeholder="IBAN" readonly={step == 2} bind:value={ payment.data.iban }>
          </div>
        </div>
      {:else if payment.method == 'CARD' }
        <div class="form-group row">
          <label class="col-sm-4 col-form-label">Your Credit Card details</label>
          <div class="col-sm-8">
            <input type="text" class="form-control" placeholder="CC number" readonly={step == 2} bind:value={ payment.data.cc_number }>
          </div>
          <div class="col-sm-8 offset-sm-4">
            <input type="text" class="form-control" placeholder="CC validity" readonly={step == 2} bind:value={ payment.data.cc_valid }>
          </div>
        </div>
      {/if}
    </div>
    <div class="col-12 d-flex justify-content-end">
      {#if step == 2}
        <button type="button" class="btn btn-info" style="margin-right: 10px" on:click={ ()=>{ progressbar.style.width = "66%"; step = 1; } }>Revert</button>
        <button type="submit" class="btn btn-success">Place an order</button>
      {:else}
        <button type="submit" class="btn btn-success">Confirm cart</button>
      {/if}
    </div>
  </form>
</div>


<script>
  import axios from 'axios';
  import of from 'await-of';
  import Auth from '../components/Auth';
  import Cart from '../components/Cart';
  import { push, replace, link } from 'svelte-spa-router'
  import { onMount } from 'svelte';

  let cart = {
    products: { data: [] }
  };
  let profile = {};
  let payment = {
    method: 'IBAN',
    data: {}
  };

  let step = 1;
  let progressbar;

  async function placeOrder() {
    if (step++ != 2) {
      progressbar.style.width = "100%"
      return;
    }
    let result = await Cart.checkout({ address: profile.address, payment: payment });
    if (result) {
      alert('Order created!');
      push('/');
    } else {
      alert('Something is wrong');
      progressbar.style.width = "66%";
      step = 1;
    }

  }

  onMount(async () => {

    if (!Auth.isLogged())
      push('/sign');

    cart = await Cart.get();
    if (cart.products.length < 1)
      replace('/cart');

    profile = await Auth.profile();

  });
</script>