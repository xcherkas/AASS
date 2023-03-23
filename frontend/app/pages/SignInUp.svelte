<div class="sign-page container">
  <div class="row">
    <div class="col-12 col-md-6">
      <h1>Sign In</h1>
      <form on:submit|preventDefault={ ()=>{ Auth.auth(signIn).then(pop) } }>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Email</label>
          <div class="col-sm-10">
            <input type="email" class="form-control" placeholder="Email" bind:value={ signIn.email }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Password</label>
          <div class="col-sm-10">
            <input type="password" class="form-control" placeholder="Password" bind:value={ signIn.password }>
          </div>
        </div>
        <div class="form-group row">
          <div class="col-sm-10">
            <button type="submit" class="btn btn-success">Sign in</button>
          </div>
        </div>
      </form>
    </div>
    <div class="col-12 col-md-6">
      <h1>Sign Up</h1>
      <form on:submit|preventDefault={ ()=>{ (signUp.password == signUp.password_confirmation)&&(Auth.register(signUp).then(pop)) } }>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Full&nbsp;name</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" placeholder="Full name" bind:value={ signUp.name }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Email</label>
          <div class="col-sm-10">
            <input type="email" class="form-control" placeholder="Email" bind:value={ signUp.email }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Phone</label>
          <div class="col-sm-10">
            <input type="text" class="form-control" placeholder="Phone number" bind:value={ signUp.phone }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label">Password</label>
          <div class="col-sm-10">
            <input type="password" class="form-control" placeholder="Password" bind:value={ signUp.password }>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-2 col-form-label"></label>
          <div class="col-sm-10">
            <input type="password" class="form-control" placeholder="Repeat your password" bind:value={ signUp.password_confirmation }>
          </div>
        </div>
        <div class="form-group row">
          <div class="col-sm-10">
            <button type="submit" class="btn btn-success">Sign up</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>


<script>
  import Auth from '../components/Auth';
  import { push, pop, replace } from 'svelte-spa-router';
  import { onMount } from 'svelte';

  let signIn = {}
  let signUp = {}

  onMount(async () => {
    if (Auth.isLogged())
      replace('/profile');
  });

</script>