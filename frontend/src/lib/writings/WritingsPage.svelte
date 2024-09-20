<script>
  import StruggleLoadingBar from "../Shared/StruggleLoadingBar.svelte";
  import TopBar from "../Shared/TopBar.svelte";
</script>

<div class='md-layout relative'>
  <TopBar />
  <div class="absolute top-2/4 left-2/4 transform -translate-x-2/4 -translate-y-2/4 w-3/4">
    <StruggleLoadingBar/>
  </div>
</div>