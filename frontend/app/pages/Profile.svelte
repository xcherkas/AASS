<div class="sign-page container">
  <div class="row">
    <div class="col-12">
      <h1>Profile of { profile.name }</h1>
      <form>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Full&nbsp;name</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" placeholder="Full name" readonly bind:value={ profile.name }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Email</label>
          <div class="col-sm-10">
            <input type="email" class="form-control" placeholder="Email" readonly bind:value={ profile.email }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Phone</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" placeholder="Phone number" readonly bind:value={ profile.phone }>
          </div>
        </div>
        <div class="form-group row">
          <div class="col-sm-10">
            <button type="button" class="btn btn-success" on:click={ logout }>Log out</button>
            {#if profile.role == 'admin' }
              <a href="/admin" class="btn btn-warning">Go to admin (max. 10 min)</a>
            {/if}
          </div>
        </div>
      </form>
    </div>
  </div>
</div>

<script>
  import axios from 'axios';
  import of from 'await-of';
  import Auth from '../components/Auth';
  import { replace } from 'svelte-spa-router'
  import { onMount } from 'svelte'

  let profile = {}

  async function changePassword() {
    let [res,] = await of(axios.post('/api/user/password-change/', {
      id: profile.id,
      password: profile.password,
      new_password: profile.password_new
    }));
    if (res)
      alert('Changed');
  }

  async function logout() {
    await Auth.deAuth();
    replace('/');
  }

  onMount(async () => {
    if (!Auth.isLogged())
      replace('/sign');
    
    profile = await Auth.profile();
  });

</script>