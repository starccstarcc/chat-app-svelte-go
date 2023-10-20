<script>
	import { api } from "../api/api";


	const weatherData = api.get({city : 'Goya'})
</script>

{#await weatherData}
	<h1>...waiting</h1>
{:then data}
{#each Object.entries(data) as [key, value], i}
    <p>{key} : {JSON.stringify(value)}</p>
{/each}

{:catch error}
<p>An error occurred!</p>
{/await}