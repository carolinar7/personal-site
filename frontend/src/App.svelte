<script lang="ts">
  import { Router, Link, Route } from "svelte-routing";
  import HomePage from "./lib/home/HomePage.svelte";
  import WritingsPage from "./lib/writings/WritingsPage.svelte";

  export let url = ""
</script>

<!-- Future Router Spot -->
<Router {url}>
  <Route path="/writings" component={WritingsPage} />
  <Route path="/" component={HomePage} />
</Router>