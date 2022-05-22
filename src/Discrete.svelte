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

		let xScale = d3.scaleBand()
			.domain(data.Points.map(d => d.X))
			.range([0, width])
            .padding(0.25);
            

		svg.append("g")
			.attr("transform", "translate(0," + height + ")")
			.call(d3.axisBottom(xScale));

		let yScale = d3.scaleLinear()
			.domain([0, d3.max(data.Points, d => d.Y)])
			.range([ height, 0 ]);

    	svg.append("g")
    		.call(d3.axisLeft(yScale));

        svg.selectAll(".bar")
            .data(data.Points)
            .enter()
            .append("rect")
            .attr("x", d => xScale(d.X))
            .attr("y", d => yScale(d.Y))
            .attr("height", d => height - yScale(d.Y))
            .attr("width", xScale.bandwidth())
            .style("fill", "#dbd0e6");

        // plot cdf data
        svg.selectAll(".bar")
            .data(data.CDF)
            .enter()
            .append("rect")
            .attr("x", d => xScale(d.X))
            .attr("y", d => yScale(d.Y))
            .attr("height", d => height - yScale(d.Y))
            .attr("width", xScale.bandwidth())
            .style("fill", "#734f96");

        // plot PDF data
        svg.selectAll(".bar")
            .data(data.PDF)
            .enter()
            .append("rect")
            .attr("x", d => xScale(d.X))
            .attr("y", d => yScale(d.Y))
            .attr("height", d => height - yScale(d.Y))
            .attr("width", xScale.bandwidth())
            .style("fill", "#734f96");

	});
</script>

<div bind:this={viz} class="chart"></div>