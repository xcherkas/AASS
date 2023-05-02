<div class="filter-group">
  <span class="title">Search by name</span>
  <input class="form-text-input" type="text" bind:value={text} >
</div>
<button class="btn btn-success" type="button" on:click={ search } style="margin-bottom: 20px">Search</button>

<div class="filter-group">
  <span class="title">Category</span>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="radio" bind:group={selectedCategory} value="mobile">
      Mobile
    </label>
  </div>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="radio" bind:group={selectedCategory} value="pc">
      PC
    </label>
  </div>
  <div class="form-check">
    <label class="form-check-label">
      <input class="form-check-input" type="radio" bind:group={selectedCategory} value="tv">
      TV
    </label>
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

  let selectedCategory = category != null ? category : undefined;
  let text = "";

  async function filter() {
    if (!selectedCategory)
      return get(products);

    const [res,] = await of(axios.get(`/v1/filter?type=${selectedCategory}`));
    get(res.data);
  };

  async function search() {
    if (text.length == 0)
      return get(products);

    const [res,] = await of(axios.get(`/v1/search?query=${text}`));
    get(res.data);
  }

  $: filter(products)


</script>