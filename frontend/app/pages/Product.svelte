<div class="container product-item">
  <div class="row">
    <div class="col-12 col-md-4 img">
      <img src="{ product.img }" alt="Coffee image">
    </div>
    <div class="col-12 col-md-7 offset-md-1 desc">
      <h2>{ product.title } &mdash; ${ product.price }</h2>
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
      <button class="btn btn-success addtocart" type="button" on:click={ ()=>{ Cart.add(product, count) } }>Add to cart</button>
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
    let [res,] = await of(axios.post(`/v1/proxy/ms-details`, {
      id: parseInt(params.id, 10)
    }));
    product = res.data;
  })();

</script>