<svelte:head>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.css" integrity="sha384-AfEj0r4/OFrOo5t7NnNe46zW/tFgW6x/bCJG8FqQCEo3+Aro6EYUG4+cU+KJWu/X" crossorigin="anonymous">
</svelte:head>

<script>
	import Katex from "svelte-katex";

	import Nav from "./Nav.svelte";
	import Continous from "./Continous.svelte";
	import Discrete from "./Discrete.svelte";
	import distributions from "./distributions.js";

	let modes = ["Distribution", "PDF", "CDF"];
	let formInfo = {
		dist: distributions[0],
		mode: modes[0],
		args: distributions[0].args,
		x: distributions[0].x,
	};
	$: formula = (" " + formInfo.dist.formula).slice(1);
	$: scales = formInfo.dist.scales 

	async function fetchData() {
		const response = await fetch("/api", {
			method: "POST",
			headers: {"accept": "application/json"},
			body: JSON.stringify({
				"Name": formInfo.dist.name,
				"Mode": formInfo.mode,
				"Args": formInfo.args.map(str => Number(str)), 
				"X": Number(formInfo.x),
			})
		});

		const json = await response.json()
		
		if (response.ok) {
			return json
		}
	}

	function handleMenuClick(event) {
		formInfo.dist = event.detail.selected;
        formInfo.args = formInfo.dist.args;
		formInfo.x = formInfo.dist.x;
        formInfo.mode = "Distribution";
        formInfo = {...formInfo};

		promise = fetchData();
	}


	
	let promise = fetchData();
</script>

<Nav on:menuClick={handleMenuClick} {formInfo} {distributions}/>

<main>
	<h1>{formInfo.dist.name}</h1>

	<div class="main-columns">
		<div>
			<!-- display formula -->
			<div>
				<Katex displayMode>{formula}</Katex>
			</div>
			<form>
				<!-- display different visualization modes -->
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
					<label for={param}>
						<Katex>{param} = {formInfo.args[i]}</Katex>
					</label>
					<input type="range" {...formInfo.dist.sliders[i]} bind:value={formInfo.args[i]} id={param} 
						on:change={() => promise = fetchData()}>
				{/each}

				<!-- only show x slider when PDF or CDF is selected -->
				{#if formInfo.mode != "Distribution"}
					<label for="x">
						<Katex>x = {formInfo.x}</Katex>
					</label>
					<input type="range" {...formInfo.dist.xSlider} bind:value={formInfo.x} id="x"
						on:change={() => promise = fetchData()}>
				{/if}
			</form>
		</div>


		<div>
			<!-- Add a discrete or continous display depending on distribution -->
			{#await promise then data}
				{#if formInfo.dist.continous === true}
					<Continous {data} {scales}/>
				{:else}
					<Discrete {data} {scales}/>
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
		height: 16px;
		border-radius: 8px;
		appearance: none;
		background-color: #dbd0e6;
		width: 100%;
		outline: none;
	}

	input[type="range"]:focus {
		outline: none;
	}

	input[type="range"]::-moz-range-thumb {
		border: none;
		border-radius: 50%;
		height: 20px;
		width: 20px;
		background-color: #734f96;
	}

	input[type="range"]::-moz-focus-inner {
		border: 0;
	}

	input[type="range"]::-webkit-slider-thumb {
		-webkit-appearance: none; 
		appearance: none;
		height: 20px;
		width: 20px;
		border-radius: 50%;
		appearance: none;
		background-color: #734f96;
		cursor: pointer;
	}

</style>