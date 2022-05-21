<script>
	import Nav from "./Nav.svelte"
	import Visualization from "./Visualization.svelte";

	let distributions = [{
		name: "Normal",
		modes: [
			{
				name: "Distribution",
				params: ["Mean", "Standard Deviation"],
				defaults: [0, 1],
			}, {
				name: "PDF", // display with point
				params: ["Mean", "Standard Deviation", "x"],
				defaults: [0, 1, 0]
			}, {
				name: "CDF", // display with filled area under curve
				params: ["Mean", "Standard Deviation", "x"],
				defaults: [0, 1, 0]
			}
		]
	}, {
		name: "Students' T",
		modes: [],
	}, {
		name: "Uniform",
		modes: [],
	}, {
		name: "Poisson",
		modes: [],
	}, {
		name: "Gamma",
		modes: [],
	}, {
		name: "Beta",
		modes: [],
	}];


	let currDist = distributions[0];
	let currMode = currDist.modes[0];
	let currArgs = [0, 1, 0];

	// $: currParams = distributions.find(d => d.name === currDist).modes.find(d => d.name === currMode) 

	async function fetchData() {
		const response = await fetch("/api", {
			method: "POST",
			headers: {"accept": "application/json"},
			body: JSON.stringify({
				"Name": currDist.name,
				"Mode": currMode.name,
				"Args": currArgs, 
			})
		});

		const json = await response.json()
		
		if (response.ok) {
			return json
		}
	}
	
	let promise = fetchData();


</script>


<Nav bind:currDist={currDist} {distributions}/>
<main>
	<h1>Normal Distribution</h1>

	<div class="main-columns">
		<div>
			<form>
				<label for="mode">Mode:</label>

				<div class="radiobtn-list">
					{#each currDist.modes as mode}
					<label for={mode.name}>
						<!-- could add onclick method to change number of parameters in inputs array -->
						<input type="radio" name="mode" id={mode.name} value={mode}
							bind:group={currMode}>
						<span>{mode.name}</span>
					</label>
					{/each}
				</div>

				{#each currMode.params as param, i}
				<label for={param}>{param}</label>
				<input type=text id={param} name={param} bind:value={currArgs[i]}><br>
				{/each}

				<input type=submit id="Submit" value="Submit" on:click|preventDefault={() => promise = fetchData()}>
			</form>
		</div>


		<div>
			{#await promise then data}
			<Visualization {data}/>
			{:catch error}
			<p>Error Message: {error.message}</p>
			{/await}
		</div>

	</div>

</main>

<style>
	main {
		padding: 2em;
	}

	h1 {
		font-weight: 400;
		text-align: center;
	}

	.main-columns {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-evenly;
	}

	.radiobtn-list {
		display: flex;
		gap: 1em;
		padding: 1em 0;
	}

	input[type=radio] {
		display: none;
	}

	input[type=radio]:checked ~ span {
		color: white;
		background-color: #734f96;
	}

	input[type=text] {
		background-color: #eee;
		outline: solid 2px #734f96;
	}

	input[type=text]:focus {
		outline: solid 3px #734f96;
	}

	input[type=submit] {
		padding: 5px 10px;
		border: 2px solid #734f96;
		border-radius: 30px;
		cursor: pointer;
	}

	input[type=submit]:focus {
		color: white;
		background-color: #734f96;
	}

	input[type=submit]:hover {
		color: white;
		background-color: #A682C9;
	}

	span {
		padding: 5px 10px;
		border: 2px solid #734f96;
		border-radius: 30px;
		cursor: pointer;
	}

	span:hover {
		color: white;
		background-color: #A682C9;
	}


	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>