<script lang="ts">
  export let href:string = ''
  export let onClick: ((e: MouseEvent) => void) | undefined = undefined
</script>

{#if href}
  <a
    {href}
    class="text-xs group-hover:flex py-1 px-2 border border-primary dark:border-white rounded-sm hover:bg-primary hover:text-white hover:dark:text-black" target="_blank"
  >
    <slot />
  </a>
{:else}
  <button
    class="text-xs group-hover:flex py-1 px-2 border border-primary dark:border-white rounded-sm hover:bg-primary hover:text-white hover:dark:text-black"
    onclick={onClick}
  >
    <slot />
  </button>
{/if}
