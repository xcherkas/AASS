<div class="category-header row align-items-stretch">
  <div class="col-12 col-md-6 img">
    <img src="/client/img/coffee_tea.jpg" alt="Coffee & Tee">
  </div>
  <div class="col-12 col-md-6 desc">
    <h2>Products</h2>
    <p>Do you know?<br>Coffee and tea are among the world’s most popular beverages, with black tea being the most sought-after variety of the later, accounting for 78% of all tea production and consumption</p>
  </div>
</div>

<div class="container-fluid">
  <div class="category-products row align-items-start">
    <aside class="col-sm-4 col-md-3">
      <Filter category={ params.category } bind:products={ products } get={ (filtered) => productsView = filtered }/>
    </aside>
    <div class="col-12 col-sm-8 col-md-9 products">
      <div class="row">

        {#each productsView.slice(page*perPage, (page+1)*perPage) as item}
          <div class="col-12 col-sm-6 col-md-4">
            <div class="product">
              <img src="{ item.img }" alt="{item.title}" on:click={ () => { push(`/product/${item.id}`) } }>
              <span class="title">{ item.title }</span>
              <span class="cost">${ item.price }</span>
              <p class="desc">{ item.desc }</p>
              <button type="button" class="btn btn-success" on:click={ () => { Cart.addMore(item) } }>Quick buy</button>
            </div>
          </div>
        {/each}

      </div>
      {#if pages.length > 1}
        <div class="col-12 d-flex justify-content-center">
          <nav aria-label="Page navigation">
            <ul class="pagination">

              <li class="page-item"><button class="page-link text-success" on:click={ ()=>{ page = Math.max(page - 1, 0); } }>Previous</button></li>

              {#each pages as index}
                <li class="page-item"><button class="page-link text-success" on:click={ ()=>{ page = index } }>{index + 1}</button></li>
              {/each}

              <li class="page-item"><button class="page-link text-success" on:click={ ()=>{ page = Math.min(page + 1, maxPage); } }>Next</button></li>

            </ul>
          </nav>
        </div>
        {/if}
    </div>
  </div>
</div>

<script>
  import axios from 'axios';
  import of from 'await-of';
  import Cart from '../components/Cart';
  import Filter from '../components/Filter';
  import { push } from 'svelte-spa-router'
  import { onMount } from 'svelte';

  export let params = {};
  const pairs = {
    'mobile': 'Mobile',
    'pc': 'PC',
    'tv': 'TV',
  }
  const perPage = 12;

  let products = [];
  let productsView = []
  let category = 'Product Catalogue';

  let page = 0;
  let pages = [ 0 ]
  let maxPage = 0;

  onMount(async () => {
    let [res,] = await of(axios.get('/v1/products'));
    res && (products = res.data);
  });

  $: {
    maxPage = Math.ceil((productsView.length / perPage)) - 1;
    pages = [
      0,
      page - 1,
      page,
      page + 1,
      maxPage
    ].filter(page => page >= 0 && page <= maxPage).filter((el, i, arr) => arr.indexOf(el) == i).sort();
  }
</script>