<script>
	import { onMount } from 'svelte';
	import * as d3 from 'd3';

	export let data;
	export let scales;

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
			.domain(scales.xScale)
			.range([0, width]);

		svg.append("g")
			.attr("transform", "translate(0," + height + ")")
			.call(d3.axisBottom(xScale));

		let yScale = d3.scaleLinear()
			.domain(scales.yScale)
			.range([ height, 0 ]);

    	svg.append("g")
    		.call(d3.axisLeft(yScale));

		let area = d3.area()
			.x(d => xScale(d.X))
			.y0(height)
			.y1(d => yScale(d.Y));

		svg.append("path")
			.datum(data.CDF)
			.attr("fill", "#dbd0e6")
			.attr("d", area)

		let line = d3.line()
			.x(d => xScale(d.X))
			.y(d => yScale(d.Y))

		svg.append("path")
			.datum(data.Curve)
			.attr("fill", "none")
			.attr("stroke", "#734f96")
			.attr("stroke-width", 2.5)
			.attr("d", line);

		if (data.PDF.length > 0) {
		svg.selectAll("mycircle")
			.data(data.PDF)
			.enter()
			.append("circle")
			.attr("cx", d => xScale(d.X))
			.attr("cy", d => yScale(d.Y))
			.attr("r", "8")
			.style("fill", "#734f96");
		}
	});
</script>

<div bind:this={viz} class="chart"></div>