<script lang="ts">
  export let title: string;
  export let description: string;
  export let date: Date;
</script>

<div class="py-4 border-y border-gray-400">
  <div class="flex justify-between">
    <div class="font-bold text-lg">{title}</div>
    <div>{date.toLocaleDateString()}</div>
  </div>
  <div>{description}</div>
</div>
