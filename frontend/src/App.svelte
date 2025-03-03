<script lang="ts">
  import { Router, Route } from 'svelte-routing';
  import HomePage from './lib/home/HomePage.svelte';
  import WritingsRouter from './lib/writings/WritingsRouter.svelte';
  import NiceTry from './NiceTry.svelte';

  export let url = '';
</script>

<!-- Future Router Spot -->
<Router {url}>
  <Route path="/writings/*">
    <WritingsRouter />
  </Route>
  <Route path="/" component={HomePage} />
  <Route path="*" component={NiceTry} />
</Router>
