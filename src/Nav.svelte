<script>
    export let currDist;
    export let distributions;
</script>

<nav class="navbar-container">
    <div class="navbar-header">
        <a href="/">DISTRIBUTION EXPLORER</a>
    </div>
    <ul class=navbar-menu-items>
        {#each distributions as distribution}
        <li class:selected={currDist.name === distribution.name} on:click={() => currDist = distribution}>
            {distribution.name}
        </li>
        {/each}
    </ul>
</nav>

<style>
    .navbar-container {
        display: flex;
        flex-direction: column;
        margin: 0.5em;
    }

    .navbar-header {
        text-align: center;
        margin: 0.25em;
        font-size: 2em;
    }

    ul {
        list-style-type: none;
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        margin: 0;
        padding: 0;
    }

    li {
        max-width: 130px;
        min-width: 60px;
        display: block;
        padding: 8px 16px;
        cursor: pointer;
    }
    li:hover {
		color: white;
		background-color: #A682C9;
	}

    .selected {
        color: white;
        background-color: #734f96;
    }

</style>