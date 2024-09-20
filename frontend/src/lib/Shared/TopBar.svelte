<script>
  import { Link } from "svelte-routing";

  const redirectItems = [
    {to: '/writings', name: 'Writings'},
  ]
  const externalLinkItems = [
    {url: 'https://www.linkedin.com/in/fredericksantiago/', name: 'Linkedin'},
    {url: 'https://github.com/carolinar7', name: 'Github'},
    {url: 'https://app.cal.com/fredericksantiago/30min', name: 'Cal'}
  ]
</script>

<header class='flex items-start justify-between flex-wrap'>
  <h1 class='font-bold text-2xl mb-3'>
    <Link to={"/"}>Frederick Santiago</Link>
  </h1>
  <ul class="flex gap-6">
    {#each redirectItems as redirectItem}
      <li class="text-gray-600 hover:text-black font-medium"><Link to={redirectItem.to}>{redirectItem.name}</Link></li>
    {/each}
    {#each externalLinkItems as externalLinkItem}
      <li class="text-gray-600 hover:text-black font-medium"><a href={externalLinkItem.url} target="_blank">{externalLinkItem.name}</a></li>
    {/each}
  </ul>
</header>