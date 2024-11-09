<script>
	import Header from "$lib/header.svelte";

    let { children } = $props();
</script>

{@render children()}

