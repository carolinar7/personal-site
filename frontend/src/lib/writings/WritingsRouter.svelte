<script lang="ts">
  import { Route, Router } from 'svelte-routing';
  import WritingsPage from './WritingsPage.svelte';
  import { writingOptions } from './writingsOptions';
  import NiceTry from '../../NiceTry.svelte';
  import NestedRouters from './pages/NestedRouters.svelte';
</script>

<Router>
  {#each writingOptions as writingOption}
    <Route path={`/${writingOption.id}`} component={NestedRouters} />
  {/each}
  <Route path={'/'} component={WritingsPage} />
  <Route path={'/*'} component={NiceTry} />
</Router>
