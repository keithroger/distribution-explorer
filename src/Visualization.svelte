<script>
	import { onMount } from 'svelte';
	import * as d3 from 'd3';
import { attr } from 'svelte/internal';
	let data = [
		{x: -5, y: 0.1},
		{x: -3, y:0.3},
		{x: 0, y: 0.4},
		{x: 2, y:0.24},
		{x: 4, y:0.1},
	]
	
	let viz;

	let margin = {top: 10, right: 30, bottom: 30, left: 60},
    width = 460 - margin.left - margin.right,
    height = 400 - margin.top - margin.bottom;

	onMount(() => {
		let svg = d3.select(viz)
			.append("svg")
    		.attr("width", width + margin.left + margin.right)
  			.attr("height", height + margin.top + margin.bottom)
			.append("g")
			.attr("transform", "translate(" + margin.left + "," + margin.top + ")");

		let xScale = d3.scaleLinear()
			.domain(d3.extent(data, d => d.x))
			.range([0, width]);

		svg.append("g")
			.attr("transform", "translate(0," + height + ")")
			.call(d3.axisBottom(xScale));

		let yScale = d3.scaleLinear()
			.domain([0, d3.max(data, d => d.y)])
			.range([ height, 0 ]);

    	svg.append("g")
    		.call(d3.axisLeft(yScale));

		let line = d3.line()
			.x(d => xScale(d.x))
			.y(d => yScale(d.y))

		svg.append("path")
			.datum(data)
			.attr("class", "line")
			.attr("d", line);

		svg.selectAll(".dot")
    		.data(dataset)
			.enter().append("circle") // Uses the enter().append() method
				.attr("class", "dot") // Assign a class for styling
				.attr("cx", function(d, i) { return xScale(i) })
				.attr("cy", function(d) { return yScale(d.y) })
				.attr("r", 5)
				.on("mouseover", function(a, b, c) { 
  				console.log(a) 
        	this.attr('class', 'focus')
			})
      		.on("mouseout", function() {  })
	});
</script>

<style>
	.line {
		fill: none;
		stroke: #ffab00;
		stroke-width: 3;
	}
  
	.overlay {
	fill: none;
	pointer-events: all;
	}
</style>


<div bind:this={viz} class="chart"></div>