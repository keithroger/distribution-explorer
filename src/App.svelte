<script>
	import Nav from "./Nav.svelte";
	import Continous from "./Continous.svelte";
	import Discrete from "./Discrete.svelte";

	import {distributions} from "./distributions.js";

	let modes = ["Distribution", "PDF", "CDF"];
	let formInfo = {
		dist: distributions[0],
		mode: modes[0],
		args: distributions[0].args,
	};
	console.log(formInfo);

	async function fetchData() {
		const response = await fetch("/api", {
			method: "POST",
			headers: {"accept": "application/json"},
			body: JSON.stringify({
				"Name": formInfo.dist.name,
				"Mode": formInfo.mode,
				"Args": formInfo.args.map(str => Number(str)), 
			})
		});

		const json = await response.json()
		
		if (response.ok) {
			return json
		}
	}
	
	let promise = fetchData();


</script>


<Nav bind:formInfo={formInfo} {distributions}/>

<main>
	<h1>Normal Distribution</h1>

	<div class="main-columns">
		<div>
			<form>
				<label for="mode">Mode:</label>

				<div class="radiobtn-list">
					{#each modes as mode}
					<label for={mode}>
						<input type="radio" name="mode" id={mode} value={mode}
							bind:group={formInfo.mode}
							on:change={() => promise = fetchData()}>
						<span>{mode}</span>
					</label>
					{/each}
				</div>

				<!-- create an input slider for every parameter -->
				{#each formInfo.dist.params as param, i}
					{#if param != "x" || formInfo.mode != "Distribution"}
						<label for={param}>{param} = {formInfo.args[i]}</label>
						<input type="range" {...formInfo.dist.sliders[i]} bind:value={formInfo.args[i]} id={param} 
							on:change={() => promise = fetchData()}>
					{/if}
				{/each}
			</form>
		</div>


		<div>
			{#await promise then data}
				{#if formInfo.dist.continous === true}
					<Continous {data}/>
				{:else}
					<Discrete {data}/>
				{/if}
			{:catch error}
				<div class="error">
					<p>Error Message: {error.message}</p>
				</div>
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
		justify-content: space-between;
		max-width: 70em;
		margin: 0 auto;
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

	input[type="range"] {
		background: transparent;
		height: 5px;
		appearance: none;
		width: 100%;
	}

	input[type="range"]:focus {
		outline: none;
	}

	input[type="range"]::-moz-range-track {
		background-color: #A682C9;
		border-radius: 5px;
		height: 5px;
	}

	input[type="range"]::-moz-range-thumb {
		border: none;
		border-radius: 50%;
		background-color: #734f96;
	}

	input[type="range"]::-moz-focus-inner {
		border: 0;
	}

	input[type="range"]::-webkit-slider-runnable-track {
		background-color: #A682C9;
		border-radius: 5px;
		height: 10px;  
	}

	/* slider thumb */
	input[type="range"]::-webkit-slider-thumb {
		-webkit-appearance: none; /* Override default look */
		appearance: none;
		margin-top: -12px; /* Centers thumb on the track */
		background-color: #734f96;
	}

</style>