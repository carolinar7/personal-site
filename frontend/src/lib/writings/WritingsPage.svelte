<script lang="ts">
  import TopBar from '../Shared/TopBar.svelte';
  import { Link } from 'svelte-routing';
  import { writingOptions } from './writingsOptions';
  import WritingOption from './WritingOption.svelte';
</script>

<svelte:head>
  <title>Frederick Santiago - Writings</title>
  <meta
    name="description"
    content="A collection of writings from Frederick Santiago; usually about software engineering."
  />
</svelte:head>

<div class="md-layout relative">
  <TopBar />
  <h1 class="pb-8">a collection of thoughts that were once just in my mind</h1>
  <main>
    <ul>
      {#each writingOptions as writingOption}
        <li>
          <Link to={writingOption.id}>
            <WritingOption
              title={writingOption.title}
              description={writingOption.description}
              date={writingOption.date}
            />
          </Link>
        </li>
      {/each}
    </ul>
  </main>
</div>
