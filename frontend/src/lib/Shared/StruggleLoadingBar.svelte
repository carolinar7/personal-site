<script>
  import { tweened } from 'svelte/motion';
  import { cubicOut } from 'svelte/easing';
  import { onMount } from 'svelte';

  let progress = tweened(0, {
    duration: 1500,
    easing: cubicOut
  });

  // Animate the progress on mount
  onMount(() => {
    progress.set(98);
  });

  $: struggling = false;

  $: if ($progress >= 95) {
    struggling = true; // Start struggling animation when close to 95%
  }
</script>

<div class="w-full h-2 bg-gray-400 rounded overflow-hidden">
  <div class="bg-black h-full rounded transition ease-out delay-500" class:wiggle={struggling}  style="width: {$progress}%"></div>
</div>

<style>
  /* Animation for struggling to reach 100% */
  @keyframes wiggle {
    0%, 100% {
      transform: translateX(0%);
    }
    25% {
      transform: translateX(-1%);
    }
    50% {
      transform: translateX(.1%);
    }
    75% {
      transform: translateX(-1%);
    }
  }

  .wiggle {
    animation: wiggle 0.5s infinite; /* Continuous left-right wiggle */
  }
</style>
