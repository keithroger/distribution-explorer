<script>
	import Nav from './Nav.svelte'
	import Visualization from './Visualization.svelte';
	
	let modes = [{
		name: 'Distributions',
		params: ['Mean', 'Standard Deviation'],
	 }, {
		 name: 'PDF',
		 params: ['Mean', 'Standard Deviation', 'x'],
	 }, {
		 name: 'CDF',
		 params: ['Mean', 'Standard Deviation', 'x'],
	 }];
	let modeSelected = 'Distributions';
	
	$: currMode = modes.find(d => d.name === modeSelected);

</script>


<Nav/>
<main>
	<h1>Normal Distribution</h1>

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

		{#each currMode.params as param}
		<label for={param}>{param}</label>
		<input type=text id={param} name={param}><br>
		{/each}

		<input type=submit id="Submit" value="Submit">
	</form>

	<Visualization/>

</main>

<style>
	main {
		padding: 2em;
	}

	h1 {
		font-weight: 400;
		text-align: center;
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