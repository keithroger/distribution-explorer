<script>
	import Nav from './Nav.svelte'
	import Visualization from './Visualization.svelte';
	
	let modes = [{
		name: 'Distribution',
		params: ['Mean', 'Standard Deviation'],
		defaults: [0, 1],
	 }, {
		name: 'PDF', // display with point
		params: ['Mean', 'Standard Deviation', 'x'],
		defaults: [0, 1, 0]
	 }, {
		name: 'CDF', // display with filled area under curve
		params: ['Mean', 'Standard Deviation', 'x'],
		defaults: [0, 1, 0]
	 }];

	let modeSelected = 'Distribution';
	let inputs = [0, 1, 0]
	let data = null
	
	$: currMode = modes.find(d => d.name === modeSelected);

	async function fetchData(event) {
		fetch("/api", {
			method: "POST",
			headers: {"accept": "application/json"},
			body: JSON.stringify({
				modeSelected,
				inputs,
			})
		}).then(res => res.json())
		.then(data => console.log(data))
		.catch((error) => {
			console.error('Error:', error);
});

	}

</script>


<Nav/>
<main>
	<h1>Normal Distribution</h1>

	<div class="main-columns">
		<div>
			<form>
				<label for="distribution">Mode:</label>

				<div class="radiobtn-list">
					{#each modes as mode}
					<label for={mode.name}>
						<input type="radio" id="{mode.name}" name="mode" value={mode.name} bind:group={modeSelected}>
						<span>{mode.name}</span>
					</label>
					{/each}
				</div>

				{#each currMode.params as param, i}
				<label for={param}>{param}</label>
				<input type=text id={param} name={param} bind:value={inputs[i]}><br>
				{/each}

				<input type=submit id="Submit" value="Submit" on:click|preventDefault={fetchData}>
			</form>
		</div>

		<!-- testing -->
		{data}

		<Visualization/>

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