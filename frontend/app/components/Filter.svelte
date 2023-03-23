<div class="filter-group">
  <span class="title">Category</span>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="checkbox" bind:group={categories} value="coffee">
      Coffee
    </label>
  </div>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="checkbox" bind:group={categories} value="tea">
      Tea
    </label>
  </div>
</div>

<div class="filter-group">
  <span class="title">PowerUps</span>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="checkbox" bind:group={powers} value="1">
      Intence
    </label>
  </div>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="checkbox" bind:group={powers} value="2">
      Zero-waste
    </label>
  </div>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="checkbox" bind:group={powers} value="3">
      Luxury
    </label>
  </div>
</div>

<div class="filter-group">
  <span class="title">Price up to { price }</span>
  <div class="form-group">
    <input type="range" class="form-control-range" min="0" max="1000" step="1" bind:value={ price }>
  </div>
</div>

<button class="btn btn-success" type="button" on:click={ filter }>Apply</button>


<script>
  import { onMount } from 'svelte';
  import axios from 'axios';
  import of from 'await-of';

  export let category;
  export let products;
  export let get;

  let categories = category != null ? [category] : [];
  let powers = [];
  let price = 1000;

  function filter() {

    let filtered = products.filter((el) => {
      try {

        if (categories.length > 0 && !categories.includes(el.categories.data.name.toLowerCase()))
          return false;
        
        let el_chars = el.characteristics.data.map(char => char.id.toString());

        if ( powers.length > 0 && (el_chars.length == 0 || !el_chars.every(cat => powers.includes(cat))) )
          return false;
        
        if (el.cost > price)
          return false;

        return true;
      } catch (ex) { console.log(ex); return false; }
    });
    get(filtered);
  };

  $: filter(products)


</script>