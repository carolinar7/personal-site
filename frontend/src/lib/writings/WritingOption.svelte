<script lang="ts">
  export let title: string;
  export let description: string;
  export let date: Date;

  function formatDate(date: Date): string {
    const months = [
      'Jan',
      'Feb',
      'Mar',
      'Apr',
      'May',
      'Jun',
      'Jul',
      'Aug',
      'Sep',
      'Oct',
      'Nov',
      'Dec'
    ];
    const day = date.getDate();
    const month = months[date.getMonth()];
    const year = date.getFullYear();
    return `${month} ${day}, ${year}`;
  }
</script>

<div class="py-4 border-y border-gray-400">
  <div class="flex justify-between">
    <div class="font-bold text-lg">{title}</div>
    <div>{formatDate(date)}</div>
  </div>
  <div>{description}</div>
</div>
