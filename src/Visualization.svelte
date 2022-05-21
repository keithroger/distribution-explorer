<script>
	import { onMount } from 'svelte';
	import * as d3 from 'd3';
	//import { attr } from 'svelte/internal';

	export let data;

	console.log(data)
	
	let viz;

	let margin = {top: 10, right: 30, bottom: 30, left: 60},
    width = 647 - margin.left - margin.right,
    height = 400 - margin.top - margin.bottom;

	onMount(() => {
		let svg = d3.select(viz)
			.append("svg")
    		.attr("width", width + margin.left + margin.right)
  			.attr("height", height + margin.top + margin.bottom)
			.append("g")
			.attr("transform", "translate(" + margin.left + "," + margin.top + ")");

		let xScale = d3.scaleLinear()
			.domain(d3.extent(data, d => d.X))
			.range([0, width]);

		svg.append("g")
			.attr("transform", "translate(0," + height + ")")
			.call(d3.axisBottom(xScale));

		let yScale = d3.scaleLinear()
			.domain([0, d3.max(data, d => d.Y)])
			.range([ height, 0 ]);

    	svg.append("g")
    		.call(d3.axisLeft(yScale));

		let line = d3.line()
			.x(d => xScale(d.X))
			.y(d => yScale(d.Y))

		svg.append("path")
			.datum(data)
			.attr("fill", "none")
			.attr("stroke", "#734f96")
			.attr("stroke-width", 2.5)
			.attr("d", line);

		// Add a lower filled section for cdf data

		// svg.selectAll(".dot")
    	// 	.data(data)
		// 	.enter().append("circle") // Uses the enter().append() method
		// 		.attr("class", "dot") // Assign a class for styling
		// 		.attr("cx", function(d, i) { return xScale(i) })
		// 		.attr("cy", function(d) { return yScale(d.y) })
		// 		.attr("r", 5)
		// 		.on("mouseover", function(a, b, c) { 
  		// 		console.log(a) 
        // 	this.attr('class', 'focus')
		// 	})
      	// 	.on("mouseout", function() {  })
	});
</script>

<div bind:this={viz} class="chart"></div>