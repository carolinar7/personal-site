<script>
  import TopBar from '../Shared/TopBar.svelte';
  import HomeDescription from './HomeDescription.svelte';
  import * as _styles from '../Shared/styles/PageLayout.css';
  // 1200 x 1600
  import meInParisImg from '../../assets/Me_in_Paris.jpg';
</script>

<svelte:head>
  <title>Frederick Santiago</title>
  <meta name="description" content="Frederick Santiago's personal website." />
</svelte:head>

<div class="md-layout">
  <TopBar />
  <main>
    <HomeDescription />
  </main>
  <img class="profileImg" src={meInParisImg} alt="Frederick in Paris" />
</div>
