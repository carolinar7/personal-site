<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import StruggleLoadingBar from "../Shared/StruggleLoadingBar.svelte";

  let { id } = $props();

  export let url = ""

  const redirectItems = [
    {to: '/abc', component: StruggleLoadingBar},
  ]
</script>

<!-- Future Router Spot -->
<Router {url}>
  {#each redirectItems as redirectedItem}
    <Route path={redirectedItem.to} component={redirectedItem.component} />
  {/each}
</Router>