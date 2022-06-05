let distributions = [
    {
        name: "Beta",
        formula: String.raw`f(x;\alpha,\beta) = \dfrac{\Gamma(\alpha + \beta)}{\Gamma(\alpha)\Gamma(\beta)}(1-x)^{\beta-1}x^{\alpha-1}`,
        continous: true,
        params: ["\\alpha", "\\beta"],
        args: ["1", "7" ],
        x: "0.5",
        scales: {
            xScale: [0, 1],
            yScale: [0, 4],
        },
        sliders: [
            {
                min: "1.0",
                max: "20.0",
                step: "0.5",
            }, {
                min: "1.0",
                max: "20.0",
                step: "0.5",
            },
        ],
        xSlider: {
            min: "0.00",
            max: "1.00",
            step: "0.05",
        },
    }, {
        name: "Binomial",
        formula: String.raw`f(x;n,p)=\binom{n}{x}p^x(1-p)^{n-x}`,
        continous: false,
        params: ["n", "p"],
        args: ["10", "0.5"],
        x: "5",
        scales: {
            xScale: [0, 24],
            yScale: [0, 1],
        },
        sliders: [
            {
                min: "1",
                max: "24",
            }, {
                min: "0.05",
                max: "0.95",
                step: "0.05",
            },
        ],
        xSlider: {
            min: "0",
            max: "24",
        },
    }, {
        name: "Chi-Squared",
        formula: String.raw`f(x; k) = \frac{e^{-x/2}x^{k/2-1}} {2^{k/2}\Gamma(k/2) }`,
        continous: true,
        params: ["k"],
        args: ["4"],
        x: "2",
        scales: {
            xScale: [0, 40],
            yScale: [0, 0.25],
        },
        sliders: [
            {
                min: "2",
                max: "20",
            },
         ],
         xSlider: {
            min: "0",
            max: "40",
        },
    }, {
        name: "Exponential",
        formula: String.raw`f(x;\lambda) = \lambda e^{-\lambda x}`,
        continous: true,
        params: ["\\lambda"],
        args: ["1"],
        x: "4",
        scales: {
            xScale: [0, 10],
            yScale: [0, 1],
        },
        sliders: [
            {
                min: "1",
                max: "10",
            },
         ],
         xSlider: {
            min: "0",
            max: "10",
        },
    }, {
        name: "Gamma",
        formula: String.raw`f(x; \alpha, \beta) = \dfrac{x^{\alpha-1}e^{-x/\beta}}{\Gamma(\alpha)\beta^\alpha}`,
        continous: true,
        params: ["\\alpha", "\\beta"],
        args: ["5", "5"],
        x: "5",
        scales: {
            xScale: [0, 25],
            yScale: [0, 1],
        },
        sliders: [
            {
                min: "1",
                max: "20",
            }, {
                min: "1",
                max: "20",
            },
         ],
        xSlider: {
            min: "0",
            max: "20",
        },
    }, {
        name: "Normal",
        formula: String.raw`f(x; \mu, \sigma) = \frac{1}{\sigma \sqrt{ 2 \pi }} \exp\left( - \frac{ (x - \mu)^2 } {2 \sigma^2}\right)`,
        continous: true,
        params: ["\\mu", "\\sigma"],
        args: ["0", "1"],
        x: "0",
        scales: {
            xScale: [-4, 4],
            yScale: [0, 0.4],
        },
        sliders: [
            {
                min: "-4",
                max: "4",
            }, {
                min: "1",
                max: "5",
                step: "0.5",
            },
         ],
        xSlider: {
            min: "-4",
            max: "4",
        },
    }, {
        name: "Poisson",
        formula: String.raw`f(x; \lambda) = \dfrac{\lambda ^x}{x!}e^{-\lambda }`,
        continous: false,
        params: ["\\lambda"],
        args: ["4"],
        x: "2",
        scales: {
            xScale: [0, 24],
            yScale: [0, 0.4],
        },
        sliders: [
            {
                min: "1",
                max: "25",
            }, 
        ],
        xSlider: {
            min: "0",
            max: "24",
        },
    }, {
        name: "Student's T",
        formula: String.raw`f(x; \hat{\mu}, \hat{\sigma}) = \frac{\Gamma((\nu+1)/2)}{\Gamma(\nu/2)\sqrt{\nu\pi\hat{\sigma}^2}} {\left(1+\dfrac{(x-\hat{\mu})^2}{\nu\hat{\sigma}^2}\right)}^{\!-(\nu+1)/2}`,
        continous: true,
        params: ["\\hat{\\mu}", "\\hat{\\sigma}", "\\nu"],
        args: ["0", "1", "3"],
        x: "0",
        scales: {
            xScale: [-4, 4],
            yScale: [0, 0.4],
        },
        sliders: [
            {
                min: "-4",
                max: "4",
            }, {
                min: "1",
                max: "5",
                step: "0.5",
            }, {
                min: "1",
                max: "50",
            },
        ],
        xSlider: {
            min: "-4",
            max: "4",
        },
    }
];

export default distributions;