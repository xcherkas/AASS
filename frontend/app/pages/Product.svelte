<div class="container product-item">
  <div class="row">
    <div class="col-12 col-md-4 img">
      <img src="{ product.img }" alt="Coffee image">
    </div>
    <div class="col-12 col-md-7 offset-md-1 desc">
      <h2>{ product.name } &mdash; ${ product.cost }</h2>
      <div class="d-flex features">
        {#each product.characteristics.data as char}
          <div class="feature">
            <img src="{ char.img }" alt="{char.name}" title="{char.name}">
            <span class="name">{ char.name }</span>
          </div>
        {/each}
      </div>
      <p>{ product.desc }</p>
      <div class="input-group col-8 col-sm-4 col-md-3 count">
        <div class="input-group-prepend">
          <button class="btn btn-success" type="button" on:click={ ()=>{ count = Math.max(1, count - 1); } }>-</button>
        </div>
        <input type="text" class="form-control" placeholder="Count" bind:value={ count }>
        <div class="input-group-append">
          <button class="btn btn-success" type="button" on:click={ ()=>{ count++; } }>+</button>
        </div>
      </div>
      <button class="btn btn-success addtocart" type="button" on:click={ ()=>{ Cart.add(product.id, count) } }>Set to cart</button>
    </div>
  </div>
</div>


<script>
  import axios from 'axios';
  import of from 'await-of';
  import Cart from '../components/Cart';


  export let params = {};
  let count = 1;
  let product = {
    characteristics: { data: [] }
  };

  (async ()=>{
    let [res,] = await of(axios.post('/api/product/get-one', {
      id: params.id
    }));
    product = res.data.data;

    let img = '//placehold.it/200';
    if (product.images.data[0])
      product.img = product.images.data[0].link;
  })();
  
</script>