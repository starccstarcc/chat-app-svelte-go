<script></script>

<svelte:head>
	<title>Svelte + GO + VIM</title>
	<meta name="description" content="First Svelte APP + Trying VIM Montions" />
</svelte:head>

<!-- <Weather/> -->

<main class="container">
	<div class="calculator">
		<div class="results" />
		<div class="digits">
			<div class="numbers" />
			<button>C</button>
			<button>1</button>
			<button>2</button>
			<button>3</button>
			<button>4</button>
			<button>5</button>
			<button>6</button>
			<button>7</button>
			<button>8</button>
			<button>9</button>
			<button>0</button>
			<button>.</button>
		</div>
		<div class="operations">
			<button>/</button>
			<button>X</button>
			<button>+</button>
			<button>-</button>
			<button>=</button>
		</div>
	</div>
</main>

<style>
	.container {
		width: 100vw;
		height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.calculator {
		display: flex;
		background-color: rgb(28, 28, 28);
		width: 400px;
		padding: 20;
		border-radius: 10px;
	}

	.digits {
		display: flex;
	}

	.operations {
		display: flex;
		flex-direction: column;
	}

	.numbers {
		display: flex;
		flex-wrap: wrap;
		width: 200px;
	}

	.results {
		display: flex;
		flex-direction: row-reverse;
		color: white;
		font-size: 40px;
		margin-right: 10px;
		height: 60px;
	}
</style>
