<script>
  let linkItems = [
    {url: 'https://www.linkedin.com/in/fredericksantiago/', name: 'Linkedin'},
    {url: 'https://github.com/carolinar7', name: 'Github'},
    {url: 'https://app.cal.com/fredericksantiago/30min', name: 'Cal'}
  ]
</script>

<header class='flex items-start justify-between flex-wrap'>
  <h1 class='font-bold text-2xl mb-3'>
    <a href='/'>Frederick Santiago</a>
  </h1>
  <ul class="flex gap-6">
    {#each linkItems as linkItem}
      <li class="text-gray-600 hover:text-black font-medium"><a href={linkItem.url} target="_blank">{linkItem.name}</a></li>
    {/each}
  </ul>
</header>