<script>
	import { api } from "../api/api";
	import { formatJSON, formatValuesIntoString } from "../utils/json";


	const weatherData = api.get({city : 'Goya'})
</script>

{#await weatherData}
	<h1>...waiting</h1>
{:then data}
{#each Object.values(data) as value}
    <p>{formatValuesIntoString(value)}</p>
{/each}

{:catch error}
<p>An error occurred!</p>
{/await}