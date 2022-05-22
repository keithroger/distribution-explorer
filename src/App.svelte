<script>
	import Nav from "./Nav.svelte";
	import Continous from "./Continous.svelte";
	import Discrete from "./Discrete.svelte";

	import {distributions} from "./distributions.js";

	let formInfo = {
		dist: distributions[0],
		mode: distributions[0].modes[0],
		args: ["0", "1", "0", "0"],
	}

	async function fetchData() {
		const response = await fetch("/api", {
			method: "POST",
			headers: {"accept": "application/json"},
			body: JSON.stringify({
				"Name": formInfo.dist.name,
				"Mode": formInfo.mode.name,
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
					{#each formInfo.dist.modes as mode}
					<label for={mode.name}>
						<input type="radio" name="mode" id={mode.name} value={mode}
							bind:group={formInfo.mode}>
						<span>{mode.name}</span>
					</label>
					{/each}
				</div>

				<!-- todo: make input into sliders -->
				{#each formInfo.mode.params as param, i}
				<label for={param}>{param}</label>
				<input type=text id={param} name={param} bind:value={formInfo.args[i]}><br>
				{/each}

				<input type=submit id="Submit" value="Submit" on:click|preventDefault={() => promise = fetchData()}>
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